package rest_util

import (
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"strings"
	"time"
)

const (
	accessKey = "accessKey"
)

type Jwt struct {
	tokenKey string
}

func NewJwt(config config.Config) *Jwt {
	return &Jwt{
		tokenKey: config.App.JwtTokenKey,
	}
}

func (j *Jwt) GenerateJwtToken(payload model.TokenPayload) (generatedToken model.TokenGenerated, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       payload.Expiration.Unix(),
		"accessKey": payload.AccessKey,
	})

	token, err := t.SignedString([]byte(payload.TokenKey))

	if err != nil {
		slog.Error(fmt.Sprint("Error from jwt.GenerateJwtToken : ", err.Error()))
		return
	}

	return model.TokenGenerated{
		Token:     token,
		ExpiredAt: payload.Expiration,
	}, nil
}

func (j *Jwt) parseJwtToken(tokenKey string, jwtToken string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error(fmt.Sprint("Unexpected signing method: ", t.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenKey), nil
	})
}

// Verify verifies the jwt token against the secret
func (j *Jwt) reverseJwtToken(tokenKey string, jwtToken string) (*model.TokenReversed, error) {
	var expiredAt time.Time
	parsed, err := j.parseJwtToken(tokenKey, jwtToken)

	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			slog.Error(fmt.Sprint("Error from jwt.parseJwtToken : ", err.Error()))
		}
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		slog.Error("Error from jwt.reverseJwtToken on parsed.Claims")
		return nil, err
	}

	accessKey, ok := claims["accessKey"].(string)
	if !ok {
		slog.Error("Error from jwt.reverseJwtToken when parsing accessKey")
		return nil, errorsConst.ErrInternalServer
	}

	expiredAtFloat, ok := claims["exp"].(float64)
	if ok {
		seconds := int64(expiredAtFloat)
		expiredAt = time.Unix(seconds, 0)
	} else {
		slog.Error("Error from jwt.reverseJwtToken when parsing exp")
		return nil, errorsConst.ErrInternalServer
	}

	return &model.TokenReversed{
		AccessKey: accessKey,
		ExpiredAt: expiredAt,
	}, nil
}

func (j *Jwt) JwtAuthenticate(ctx *fiber.Ctx) error {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		return errorsConst.ErrUnauthorized
	}

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		return errorsConst.ErrUnauthorized
	}

	// Spliting the header
	chunks := strings.Split(authorization, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return errorsConst.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	reversedToken, err := j.reverseJwtToken(j.tokenKey, chunks[1])
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			slog.Error(fmt.Sprintln("Error on reverse jwt token : ", err.Error()))
		}
		return errorsConst.ErrUnauthorized
	}

	ctx.Locals(accessKey, reversedToken.AccessKey)
	return ctx.Next()
}

func GetJwtAccessKey(ctx *fiber.Ctx) string {
	return ctx.Locals(accessKey).(string)
}

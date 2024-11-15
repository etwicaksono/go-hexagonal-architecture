package rest_util

import (
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/cache"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"strings"
	"time"
)

const (
	accessKey        = "accessKey"
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

type Jwt struct {
	tokenKey        string
	tokenExpiration string
	tokenRefresh    string
	cache           cache.CacheInterface
}

func NewJwt(
	config config.Config,
	cache cache.CacheInterface,
) *Jwt {
	return &Jwt{
		tokenKey:        config.App.JwtTokenKey,
		tokenExpiration: config.App.JwtTokenExpiration,
		tokenRefresh:    config.App.JwtTokenRefresh,
		cache:           cache,
	}
}

func (j *Jwt) GenerateJwtToken(accessKey string) (generatedToken model.TokenGenerated, err error) {
	// Generate the access token
	accessTokenAdditionalDuration, err := time.ParseDuration(j.tokenExpiration)
	if err != nil {
		return
	}
	expiredAt := time.Now().Add(accessTokenAdditionalDuration)
	accessKeyClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       expiredAt.Unix(),
		"iat":       time.Now().Unix(),
		"accessKey": accessKey,
		"type":      AccessTokenType,
	})
	accessToken, err := accessKeyClaims.SignedString([]byte(j.tokenKey))

	// Generate the refresh token
	refreshTokenAdditionalDuration, err := time.ParseDuration(j.tokenRefresh)
	if err != nil {
		return
	}
	refreshableUntil := time.Now().Add(refreshTokenAdditionalDuration)
	refreshKeyClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       refreshableUntil.Unix(),
		"iat":       time.Now().Unix(),
		"accessKey": accessKey,
		"type":      RefreshTokenType,
	})
	refreshToken, err := refreshKeyClaims.SignedString([]byte(j.tokenKey))

	if err != nil {
		slog.Error(fmt.Sprint("Error from jwt.GenerateJwtToken : ", err.Error()))
		return
	}

	return model.TokenGenerated{
		AccessToken:      accessToken,
		ExpiredAt:        expiredAt,
		RefreshToken:     refreshToken,
		RefreshableUntil: refreshableUntil,
	}, nil
}

func (j *Jwt) parseJwtToken(jwtToken string) (*jwt.Token, error) {
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
		return []byte(j.tokenKey), nil
	})
}

// Verify verifies the jwt token against the secret
func (j *Jwt) ReverseJwtToken(jwtToken string) (*entity.AuthToken, error) {
	var expiredAt time.Time
	parsed, err := j.parseJwtToken(jwtToken)

	if err != nil {
		slog.Error(fmt.Sprint("Error from jwt.parseJwtToken : ", err.Error()))
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		slog.Error("Error from jwt.ReverseJwtToken on parsed.Claims")
		return nil, err
	}

	typeParsed, ok := claims["type"].(string)
	if !ok {
		slog.Error("Error from jwt.ReverseJwtToken when parsing accessKey")
		return nil, errorsConst.ErrInternalServer
	}

	accessKeyParsed, ok := claims["accessKey"].(string)
	if !ok {
		slog.Error("Error from jwt.ReverseJwtToken when parsing accessKey")
		return nil, errorsConst.ErrInternalServer
	}

	expiredAtFloat, ok := claims["exp"].(float64)
	if ok {
		seconds := int64(expiredAtFloat)
		expiredAt = time.Unix(seconds, 0)
	} else {
		slog.Error("Error from jwt.ReverseJwtToken when parsing exp")
		return nil, errorsConst.ErrInternalServer
	}

	return &entity.AuthToken{
		AccessKey: accessKeyParsed,
		ExpiredAt: expiredAt,
		TokenType: typeParsed,
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
	reversedToken, err := j.ReverseJwtToken(chunks[1])
	if err != nil {
		slog.Error(fmt.Sprintln("Error on reverse jwt token : ", err.Error()))
		return errorsConst.ErrUnauthorized
	}

	if reversedToken.TokenType != AccessTokenType {
		return errorsConst.ErrUnauthorized
	}

	// Check if token is within blacklist
	_, err = j.cache.GetAuthToken(ctx.UserContext(), fmt.Sprintf("%s:%s", constants.BlackListedTokenRedisPrefix, reversedToken.AccessKey))
	if err == nil {
		return errorsConst.ErrUnauthorized
	} else {
		if !errors.Is(redis.Nil, err) {
			slog.ErrorContext(ctx.UserContext(), "Failed to get auth token from redis", slog.String("error", err.Error()))
		}
	}

	ctx.Locals(accessKey, reversedToken)
	return ctx.Next()
}

func (j *Jwt) GetJwtAuthToken(ctx *fiber.Ctx) (authToken *entity.AuthToken, err error) {
	authToken, ok := ctx.Locals(accessKey).(*entity.AuthToken)
	if !ok {
		return authToken, errorsConst.ErrTokenClaimsParseFailed
	}
	return
}

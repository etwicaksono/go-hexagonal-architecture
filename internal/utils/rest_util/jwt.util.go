package rest_util

import (
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

func Generate(payload model.TokenPayload) model.TokenGenerated {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       payload.Expiration.Unix(),
		"accessKey": payload.AccessKey,
	})

	token, err := t.SignedString([]byte(payload.TokenKey))

	if err != nil {
		slog.Error(fmt.Sprint("Error from jwt.Generate : ", err.Error()))
		panic(err)
	}

	return model.TokenGenerated{
		Token:     token,
		ExpiredAt: payload.Expiration,
	}
}

func parse(tokenKey string, jwtToken string) (*jwt.Token, error) {
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
func Reverse(tokenKey string, jwtToken string) (*model.TokenReversed, error) {
	var expiredAt time.Time
	parsed, err := parse(tokenKey, jwtToken)

	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			slog.Error(fmt.Sprint("Error from jwt.parse : ", err.Error()))
		}
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		slog.Error("Error from jwt.Reverse on parsed.Claims")
		return nil, err
	}

	accessKey, ok := claims["accessKey"].(string)
	if !ok {
		slog.Error("Error from jwt.Reverse when parsing accessKey")
		return nil, errors2.ErrInternalServer
	}

	expiredAtFloat, ok := claims["exp"].(float64)
	if ok {
		seconds := int64(expiredAtFloat)
		expiredAt = time.Unix(seconds, 0)
	} else {
		slog.Error("Error from jwt.Reverse when parsing exp")
		return nil, errors2.ErrInternalServer
	}

	return &model.TokenReversed{
		AccessKey: accessKey,
		ExpiredAt: expiredAt,
	}, nil
}

package authentication_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"log/slog"
	"time"
)

func (a authenticationCore) Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error) {
	// check if user exist
	user, err := a.db.FindByFilter(ctx, entity.UserFindFilter{Email: request.Email})
	if err != nil {
		slog.ErrorContext(ctx, "Error on finding user", slog.String(entity.Error, err.Error()))
		if errors.Is(err, errors2.ErrNoData) {
			return entity.TokenGenerated{}, errors2.ErrUserNotFound
		}
		return
	}

	// verify password
	if err = utils.PasswordVerify(user.Password, request.Password); err != nil {
		return entity.TokenGenerated{}, errors2.ErrInvalidLoginCredentials
	}

	accessKey, err := utils.PasswordGenerate(user.ID)
	if err != nil {
		return
	}

	additionalDuration, err := time.ParseDuration(a.config.App.JwtExpiration)
	if err != nil {
		return
	}
	expiredAt := time.Now().Add(additionalDuration)
	generatedJwt, err := rest_util.Generate(model.TokenPayload{
		AccessKey:  accessKey,
		TokenKey:   a.config.App.JwtTokenKey,
		Expiration: expiredAt,
	})

	if err != nil {
		return
	}

	return generatedJwt.ToEntity(), nil
}

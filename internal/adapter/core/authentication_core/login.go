package authentication_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils"
	"github.com/guregu/null"
	"log/slog"
)

func (a authenticationCore) Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error) {
	// check if user exist
	user, err := a.db.FindByFilter(ctx, entity.UserFindFilter{Email: null.StringFrom(request.Email)})
	if err != nil {
		slog.ErrorContext(ctx, "Error on finding user", slog.String(entity.Error, err.Error()))
		if errors.Is(err, errors2.ErrNoData) {
			return entity.TokenGenerated{}, errors2.ErrUserNotFound
		}
		return
	}

	// verify password
	if err = utils.PasswordVerify(user.Password, request.Password); err != nil {
		return entity.TokenGenerated{}, errors2.ErrLoginCredentialInvalid
	}

	accessKey, err := utils.PasswordGenerate(user.ID)
	if err != nil {
		return
	}
	generatedJwt, err := a.jwt.GenerateJwtToken(accessKey)
	if err != nil {
		return
	}

	// TODO: save access key to cache

	return generatedJwt.ToEntity(), nil
}

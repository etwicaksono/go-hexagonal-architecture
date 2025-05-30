package authentication_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils"
	"github.com/guregu/null"
	"log/slog"
)

func (a authenticationCore) Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error) {
	// check if user exist
	user, err := a.db.FindByFilter(ctx, entity.UserFindFilter{Email: null.StringFrom(request.Email)})
	if err != nil {
		if errors.Is(err, errorsConst.ErrNoData) {
			return entity.TokenGenerated{}, errorsConst.ErrUserNotFound
		}
		slog.ErrorContext(ctx, "Error on finding user", slog.String("email", request.Email), slog.String(constants.Error, err.Error()))
		return
	}

	// verify password
	if err = utils.PasswordVerify(user.Password, request.Password); err != nil {
		return entity.TokenGenerated{}, errorsConst.ErrLoginCredentialInvalid
	}

	accessKey, err := utils.PasswordGenerate(user.ID)
	if err != nil {
		slog.ErrorContext(ctx, "Error on generating access key", slog.String(constants.Error, err.Error()))
		return
	}
	generatedJwt, err := a.jwt.GenerateJwtToken(accessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on generating jwt token", slog.String(constants.Error, err.Error()))
		return
	}

	err = a.cache.SetAuthenticatedToken(ctx, accessKey, model.AuthCachedData{
		UserId:    user.ID,
		AccessKey: accessKey,
		ExpiredAt: generatedJwt.ExpiredAt,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error on saving auth token to cache", slog.String(constants.Error, err.Error()))
		return
	}

	return generatedJwt.ToEntity(), nil
}

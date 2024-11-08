package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"log/slog"
)

func (a authenticationApp) Login(ctx context.Context, request entity.LoginRequest) (result entity.TokenGenerated, err error) {
	err = validation_util.ValidateStruct(a.validator, model.FromLoginRequestEntity(request))
	if err != nil {
		return
	}

	result, err = a.core.Login(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on authenticating user", slog.String(entity.Error, err.Error()))
		return
	}

	return
}

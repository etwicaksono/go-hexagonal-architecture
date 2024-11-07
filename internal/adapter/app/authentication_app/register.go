package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"log/slog"
)

func (a authenticationApp) Register(ctx context.Context, request entity.RegisterRequest) (err error) {
	err = validation_util.ValidateStruct(a.validator, model.FromRegisterRequestEntity(request))
	if err != nil {
		return
	}

	err = a.core.Register(ctx, request)
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(ctx, "Error on registering user", slog.String(entity.Error, err.Error()))
		}
		return
	}

	return
}

package example_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"log/slog"
)

func (e exampleApp) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error {
	err := e.validator.Struct(model.FromSendMultimediaMessageRequestEntity(request))
	if err != nil {
		errValidation := utils.GenerateErrorMessage(err)
		return error_util.ValidationError(errValidation)
	}

	err = e.core.SendMultimediaMessage(ctx, request)
	if err != nil {
		if !error_util.IsValidationError(err) {
			slog.ErrorContext(ctx, "Error on sending multimedia message", slog.String(entity.Error, err.Error()))
		}
		return err
	}

	return nil
}

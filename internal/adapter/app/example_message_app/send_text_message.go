package example_message_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/payload_util"
	"log/slog"
)

func (e exampleMessageApp) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error {
	err := e.validator.Struct(model.FromSendTextMessageRequestEntity(request))
	if err != nil {
		errValidation := payload_util.GenerateErrorMessage(err)
		return error_util.ValidationError(errValidation)
	}

	err = e.core.SendTextMessage(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on sending text message", slog.String(entity.Error, err.Error()))
		return err
	}

	return nil
}

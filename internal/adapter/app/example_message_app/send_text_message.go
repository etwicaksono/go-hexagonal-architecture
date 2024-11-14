package example_message_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"log/slog"
)

func (e exampleMessageApp) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) (err error) {
	err = validation_util.ValidateStruct(e.validator, model.FromSendTextMessageRequestEntity(request))
	if err != nil {
		return
	}

	err = e.core.SendTextMessage(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on sending text message", slog.String(constants.Error, err.Error()))
		return err
	}

	return nil
}

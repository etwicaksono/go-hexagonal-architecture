package example_message_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"log/slog"
)

func (e exampleMessageApp) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) (err error) {
	err = validation_util.ValidateStruct(e.validator, model.FromSendMultimediaMessageRequestEntity(request))
	if err != nil {
		return
	}

	err = e.core.SendMultimediaMessage(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on sending multimedia message", slog.String(constants.Error, err.Error()))
		return err
	}

	return nil
}

package example_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (e exampleApp) SendTextMessage(ctx context.Context, request entity.SendTextMessageRequest) error {

	err := e.validator.Struct(model.FromSendTextMessageRequestEntity(request))
	if err != nil {
		errValidation := utils.GenerateErrorMessage(err)
		return utils.NewCustomError().
			SetCode(fiber.StatusBadRequest).
			SetMessage(utils.ValidationError).
			SetFields(errValidation)
	}

	err = e.core.SendTextMessage(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Error on sending text message", slog.String("error", err.Error()))
		return err
	}

	return nil
}

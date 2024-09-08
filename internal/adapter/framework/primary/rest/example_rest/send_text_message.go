package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a adapter) SendTextMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)

	payload := new(model.SendTextMessageRequest)
	err = ctx.BodyParser(payload)
	if err != nil {
		errParsing, errOther := utils.HandleParsingError(err)
		if errOther != nil {
			slog.ErrorContext(context, errOther.Error())
			return errOther
		}
		return error_util.ValidationError(errParsing)
	}

	err = a.app.SendTextMessage(context, payload.ToEntity())
	if err != nil {
		slog.ErrorContext(context, "Failed to send text message", slog.String(entity.Error, err.Error()))
		return err
	}

	return rest_util.ResponseOk(ctx, "Send text message success")
}
package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/gofiber/fiber/v2"
)

func (a adapter) SendTextMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)

	payload := new(model.SendTextMessageRequest)
	err = ctx.BodyParser(payload)
	if err != nil {
		errParsing, errOther := utils.HandleParsingError(err)
		if errOther != nil {
			return errOther
		}
		return error_util.NewCustomError().
			SetCode(fiber.StatusBadRequest).
			SetMessage(entity.Error).
			SetFields(errParsing)
	}

	err = a.app.SendTextMessage(context, payload.ToEntity())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response[[]model.MessageTextItem]{
		Code:    fiber.StatusOK,
		Status:  entity.Success,
		Message: "Send text message success",
	})
}

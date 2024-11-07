package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a AuthenticationHandler) Register(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.RegisterRequest)
	err = ctx.BodyParser(payload) // TODO: create util for this
	if err != nil {
		errParsing, errOther := payload_util.HandleParsingError(err)
		if errOther != nil {
			slog.ErrorContext(context, errOther.Error())
			return errOther
		}
		return error_util.ErrorValidation(errParsing)
	}

	err = a.app.Register(context, payload.ToEntity())
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to register user", slog.String(entity.Error, err.Error()))
		}
		return err
	}

	return rest_util.ResponseOk(ctx, "Register user success")
}

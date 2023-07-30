package configs

import (
	"github.com/Firhan384/restfull-api-golang/app/schemas"
	"github.com/Firhan384/restfull-api-golang/app/utils"
	"github.com/gofiber/fiber/v2"
)

func NewConfigFiber(cfg ConfigApp) fiber.Config {
	return fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			resp := schemas.NewResponseApi()

			if err == fiber.ErrInternalServerError {
				resp.StatusCode = fiber.StatusInternalServerError
				resp.Message = utils.StatusText(fiber.StatusInternalServerError)
				resp.Error = true

				return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
			}
			return nil
		},
	}
}

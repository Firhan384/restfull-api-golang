package configs

import (
	"github.com/gofiber/fiber/v2"
)

func NewConfigFiber(cfg ConfigApp) fiber.Config {
	return fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
	}
}

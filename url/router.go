package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	umkapi := app.Group("/lapumk", logger.New())
	SetuplapRoutes(umkapi)
}

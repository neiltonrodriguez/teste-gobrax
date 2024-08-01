package main

import (
	"teste-gobrax/config"
	RouterDriver "teste-gobrax/internal/driver/router"
	RouterVehicle "teste-gobrax/internal/vehicle/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// log.Info().Msg("Server started")
	app := fiber.New()
	// common.NewLogger()
	config.GlobalConfig.LoadVariables()

	RouterDriver.RegisterRoutes(app)
	RouterVehicle.RegisterRoutes(app)
	app.Listen(":8080")
}

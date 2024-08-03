package router

import (
	"teste-gobrax/internal/vehicle/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,

) {

	route := app.Group("v1/vehicle")

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Create(fiberCtx)
	})

	route.Get("/", func(fiberCtx *fiber.Ctx) error {
		return handler.GetAll(fiberCtx)
	})

	route.Get("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.GetById(fiberCtx)
	})

	route.Put("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Update(fiberCtx)
	})

	route.Delete("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Delete(fiberCtx)
	})

}

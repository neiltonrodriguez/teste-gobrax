package handler

import (
	"strconv"
	domain "teste-gobrax/internal/domain"
	Model "teste-gobrax/internal/driver/model"

	"github.com/gofiber/fiber/v2"
)

func Get(fiberCtx *fiber.Ctx) error {

	ctx := fiberCtx.Context()

	result, err := Model.Get(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: len(result),
		},
		Data: result,
	})
}

func GetById(fiberCtx *fiber.Ctx) error {
	ctx := fiberCtx.Context()
	id, _ := strconv.Atoi(fiberCtx.Params("id"))

	result, err := Model.GetById(ctx, id)
	if err != nil {
		if err.Error() == "user not found" {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: result,
	})
}

func Create(fiberCtx *fiber.Ctx) error {
	payload := domain.Driver{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := Model.Create(ctx, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: result,
	})
}

func Update(fiberCtx *fiber.Ctx) error {
	payload := domain.Driver{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(fiberCtx.Params("id"))
	err := Model.Update(ctx, id, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "updated with success",
	})
}

func Delete(fiberCtx *fiber.Ctx) error {
	ctx := fiberCtx.Context()

	id, _ := strconv.Atoi(fiberCtx.Params("id"))

	err := Model.Delete(ctx, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "Deleted with success",
	})
}

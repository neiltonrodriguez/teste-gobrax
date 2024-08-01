package handler

import (
	"strconv"
	domain "teste-gobrax/internal/domain"
	Model "teste-gobrax/internal/vehicle/model"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Get(fiberCtx *fiber.Ctx) error {

	ctx := fiberCtx.Context()

	result, err := Model.Get(ctx)
	if err != nil {
		log.Error().Msg("Error in GetAll of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("GetAll of vehicle successfully")

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
		if err.Error() == "vehicle not found" {
			log.Error().Msg("Error in GetByID of vehicle: " + err.Error())
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("GetById of vehicle successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: result,
	})
}

func Create(fiberCtx *fiber.Ctx) error {
	payload := domain.Vehicle{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("error in parse payload of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Create of vehicle successfully")

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
	payload := domain.Vehicle{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("Parse error in update of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, _ := strconv.Atoi(fiberCtx.Params("id"))
	err := Model.Update(ctx, id, payload)
	if err != nil {
		log.Error().Msg("Error in update of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Update of vehicle successfully")

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
		log.Error().Msg("Error in delete of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Delete of vehicle successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "Deleted with success",
	})
}

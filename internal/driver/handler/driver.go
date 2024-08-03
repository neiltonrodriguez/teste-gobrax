package handler

import (
	"strconv"
	domain "teste-gobrax/internal/domain"
	Model "teste-gobrax/internal/driver/model"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAll(fiberCtx *fiber.Ctx) error {

	ctx := fiberCtx.Context()
	filter := map[string]string{
		"available": fiberCtx.Query("available", "false"),
		"name":      fiberCtx.Query("name", ""),
	}

	page, _ := strconv.Atoi(fiberCtx.Query("page", "0"))
	limit, _ := strconv.Atoi(fiberCtx.Query("limit", "0"))
	pg := domain.NewPagination(page, limit)

	results, err := Model.GetAll(ctx, filter, pg)
	if err != nil {
		log.Error().Msg("Error in GetAll of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var driversDTO []domain.DriverDTO
	for _, driver := range results {
		driverDTO := domain.ConvertToDriverDTO(driver)
		driversDTO = append(driversDTO, driverDTO)
	}

	log.Info().Msg("GetAll of drivers successfully")
	total := Model.GetTotal(ctx, filter)

	if pg.Valid() {
		pg.SetTotal(total)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count:      len(driversDTO),
			Pagination: pg.ToResponse(),
		},
		Data: driversDTO,
	})
}

func GetById(fiberCtx *fiber.Ctx) error {
	ctx := fiberCtx.Context()
	id, _ := strconv.Atoi(fiberCtx.Params("id"))

	result, err := Model.GetById(ctx, id)
	if err != nil {
		if err.Error() == "driver not found" {
			log.Error().Msg("Error in GetByID of drivers: " + err.Error())
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		log.Error().Msg("Error in GetByID of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	driver := domain.ConvertToDriverDTO(result)

	log.Info().Msg("GetById of drivers successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: driver,
	})
}

func Create(fiberCtx *fiber.Ctx) error {
	payload := domain.DriverInput{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("error in parse payload of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	driveInput := domain.Driver{
		Name:           payload.Name,
		DriversLicense: payload.DriversLicense,
		Phone:          payload.Phone,
		Age:            payload.Age,
	}

	result, err := Model.Create(ctx, driveInput)
	if err != nil {
		log.Error().Msg("error in create driver: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	driver := domain.ConvertToDriverDTO(result)

	log.Info().Msg("Create of drivers successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: driver,
	})
}

func Update(fiberCtx *fiber.Ctx) error {
	payload := domain.DriverInput{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("Parse error in update of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	driveInput := domain.Driver{
		Name:           payload.Name,
		DriversLicense: payload.DriversLicense,
		Phone:          payload.Phone,
		Age:            payload.Age,
	}

	id, _ := strconv.Atoi(fiberCtx.Params("id"))
	err := Model.Update(ctx, id, driveInput)
	if err != nil {
		log.Error().Msg("Error in update of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Update of drivers successfully")

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
		log.Error().Msg("Error in delete of drivers: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Info().Msg("Delete of drivers successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: "Deleted with success",
	})
}

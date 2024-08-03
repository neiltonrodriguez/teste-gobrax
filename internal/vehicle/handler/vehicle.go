package handler

import (
	"strconv"
	domain "teste-gobrax/internal/domain"
	DriverModel "teste-gobrax/internal/driver/model"
	Model "teste-gobrax/internal/vehicle/model"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetAll(fiberCtx *fiber.Ctx) error {

	ctx := fiberCtx.Context()
	filter := map[string]string{
		"plate": fiberCtx.Query("plate", ""),
		"brand": fiberCtx.Query("brand", ""),
	}

	page, _ := strconv.Atoi(fiberCtx.Query("page", "0"))
	limit, _ := strconv.Atoi(fiberCtx.Query("limit", "0"))
	pg := domain.NewPagination(page, limit)

	vehicles, err := Model.GetAll(ctx, filter, pg)
	if err != nil {
		log.Error().Msg("Error in GetAll of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var vehiclesDTO []domain.VehicleDTO
	for _, vehicle := range vehicles {
		var vehicleDTO domain.VehicleDTO
		if vehicle.DriverId != 0 {
			driver, err := DriverModel.GetById(ctx, vehicle.DriverId)
			if err != nil {
				log.Error().Msg("Error in GetAll of vehicle: " + err.Error())
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			// driverDTO := domain.ConvertToDriverDTO(driver)
			vehicleDTO = domain.ConvertToVehicleDTO(vehicle, &driver)
		} else {
			vehicleDTO = domain.ConvertToVehicleDTO(vehicle, nil)
		}

		vehiclesDTO = append(vehiclesDTO, vehicleDTO)

	}

	log.Info().Msg("GetAll of vehicle successfully")
	total := Model.GetTotal(ctx, filter)

	if pg.Valid() {
		pg.SetTotal(total)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count:      len(vehiclesDTO),
			Pagination: pg.ToResponse(),
		},
		Data: vehiclesDTO,
	})
}

func GetById(fiberCtx *fiber.Ctx) error {
	ctx := fiberCtx.Context()
	id, _ := strconv.Atoi(fiberCtx.Params("id"))

	vehicle, err := Model.GetById(ctx, id)
	if err != nil {
		if err.Error() == "vehicle not found" {
			log.Error().Msg("Error in GetByID of vehicle: " + err.Error())
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var vehicleDTO domain.VehicleDTO
	if vehicle.DriverId != 0 {
		driver, err := DriverModel.GetById(ctx, vehicle.DriverId)
		if err != nil {
			log.Error().Msg("Error in GetAll of vehicle: " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		vehicleDTO = domain.ConvertToVehicleDTO(vehicle, &driver)
	} else {
		vehicleDTO = domain.ConvertToVehicleDTO(vehicle, nil)
	}

	log.Info().Msg("GetById of vehicle successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: vehicleDTO,
	})
}

func Create(fiberCtx *fiber.Ctx) error {
	payload := domain.VehicleInput{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("error in parse payload of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	vehicleInput := domain.Vehicle{
		DriverId: payload.DriverId,
		Plate:    payload.Plate,
		Brand:    payload.Brand,
		Model:    payload.Model,
	}

	plateExist := Model.CheckPlateExist(ctx, vehicleInput.Plate, 0)
	if plateExist {
		log.Error().Msg("plate already exists")
		return fiber.NewError(fiber.StatusBadRequest, "plate already exists")
	}

	if vehicleInput.DriverId != 0 {
		driverIsAvailable := DriverModel.CheckDriverIsAvailable(ctx, vehicleInput.DriverId)
		if !driverIsAvailable {
			log.Error().Msg("driver is not available")
			return fiber.NewError(fiber.StatusBadRequest, "driver is not available")
		}
	}

	result, err := Model.Create(ctx, vehicleInput)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var vehicle domain.VehicleDTO
	if result.DriverId != 0 {
		driver, err := DriverModel.GetById(ctx, result.DriverId)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		vehicle = domain.ConvertToVehicleDTO(result, &driver)
	} else {
		vehicle = domain.ConvertToVehicleDTO(result, nil)
	}

	log.Info().Msg("Create of vehicle successfully")

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: 1,
		},
		Data: vehicle,
	})
}

func Update(fiberCtx *fiber.Ctx) error {
	payload := domain.VehicleInput{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		log.Error().Msg("Parse error in update of vehicle: " + err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	vehicleInput := domain.Vehicle{
		DriverId: payload.DriverId,
		Plate:    payload.Plate,
		Brand:    payload.Brand,
		Model:    payload.Model,
	}

	plateExist := Model.CheckPlateExist(ctx, vehicleInput.Plate, vehicleInput.Id)
	if plateExist {
		log.Error().Msg("plate already exists")
		return fiber.NewError(fiber.StatusBadRequest, "plate already exists")
	}

	if vehicleInput.DriverId != 0 {
		driverIsAvailable := DriverModel.CheckDriverIsAvailable(ctx, vehicleInput.DriverId)
		if !driverIsAvailable {
			log.Error().Msg("driver is not available")
			return fiber.NewError(fiber.StatusBadRequest, "driver is not available")
		}
	}

	id, _ := strconv.Atoi(fiberCtx.Params("id"))
	err := Model.Update(ctx, id, vehicleInput)
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

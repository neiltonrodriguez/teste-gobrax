package domain

import "time"

type Vehicle struct {
	Id        int       `json:"id"`
	DriverId  int       `json:"driver_id"`
	Plate     string    `json:"plate"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VehicleInput struct {
	DriverId int    `json:"driver_id"`
	Plate    string `json:"plate"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
}

type VehicleDTO struct {
	Id        int        `json:"id"`
	Driver    *DriverDTO `json:"driver"`
	Plate     string     `json:"plate"`
	Brand     string     `json:"brand"`
	Model     string     `json:"model"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func ConvertToVehicleDTO(vehicle Vehicle, driver *Driver) VehicleDTO {
	var driverDTO *DriverDTO
	if driver != nil {
		dto := ConvertToDriverDTO(*driver)
		driverDTO = &dto
	}

	return VehicleDTO{
		Id:        vehicle.Id,
		Driver:    driverDTO,
		Plate:     vehicle.Plate,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}

package domain

import "time"

type Driver struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	DriversLicense string    `json:"drivers_license"`
	Phone          string    `json:"phone"`
	Age            int       `json:"age"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DriverInput struct {
	Name           string    `json:"name"`
	DriversLicense string    `json:"drivers_license"`
	Phone          string    `json:"phone"`
	Age            int       `json:"age"`
}

type DriverDTO struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	DriversLicense string    `json:"drivers_license"`
	Phone          string    `json:"phone"`
	Age            int       `json:"age"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ConvertToDriverDTO(driver Driver) DriverDTO {
    return DriverDTO{
        Id:             driver.Id,
        Name:           driver.Name,
        DriversLicense: driver.DriversLicense,
        Phone:          driver.Phone,
        Age:            driver.Age,
        CreatedAt:      driver.CreatedAt,
        UpdatedAt:      driver.UpdatedAt,
    }
}
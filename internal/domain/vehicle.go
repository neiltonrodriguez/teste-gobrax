package domain

import "time"

type Vehicle struct {
	Id        int       `json:"id"`
	DriverId  int       `json:"driver_id"`
	Placa     string    `json:"placa"`
	Marca     string    `json:"marca"`
	Modelo    string    `json:"modelo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

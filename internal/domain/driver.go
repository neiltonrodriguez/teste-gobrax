package domain

import "time"

type Driver struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	DriversLicense string    `json:"drivers_license"`
	Phone          string    `json:"phone"`
	Age            int       `json:"age"`
	Address        string    `json:"address"`
	Neighborhood   string    `json:"neighborhood"`
	Zipcode        string    `json:"zipcode"`
	City           string    `json:"city"`
	UF             string    `json:"uf"`
	Number         string    `json:"number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

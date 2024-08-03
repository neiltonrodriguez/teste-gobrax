package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"teste-gobrax/internal/domain"
	"teste-gobrax/pkg/database"
)

var Db *sql.DB

var (
	errVehicleNotFound = errors.New("vehicle not found")
)

func GetAll(ctx context.Context, filter map[string]string, pg domain.Pagination) ([]domain.Vehicle, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	var args []interface{}
	conditions := ""
	if filter["plate"] != "" {
		conditions += " WHERE v.plate = ? "
		args = append(args, filter["plate"])
	}

	if filter["brand"] != "" {
		conditions += " WHERE v.brand = ? "
		args = append(args, filter["brand"])
	}

	pagination := ""
	if pg.Valid() {
		if pg.Limit() > 0 {
			pagination += " LIMIT ?"
			args = append(args, pg.Limit())
		}
		if pg.Offset() > 0 {
			pagination += " OFFSET ?"
			args = append(args, pg.Offset())
		}
	}

	query := `
	SELECT 
		v.id, 
		COALESCE(driver_id, 0),
		v.Plate, 
		v.Brand, 
		v.Model, 
		v.created_at, 
		v.updated_at 
	FROM teste_gobrax.vehicle v
	LEFT JOIN teste_gobrax.driver d ON d.id = v.driver_id` + conditions + pagination

	rows, err := Db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []domain.Vehicle
	for rows.Next() {
		var vehicle domain.Vehicle
		err := rows.Scan(
			&vehicle.Id,
			&vehicle.DriverId,
			&vehicle.Plate,
			&vehicle.Brand,
			&vehicle.Model,
			&vehicle.CreatedAt,
			&vehicle.UpdatedAt)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

func GetTotal(ctx context.Context, filter map[string]string) int {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return 0
	}

	var args []interface{}
	conditions := ""
	if filter["plate"] != "" {
		conditions += " WHERE v.plate = ? "
		args = append(args, filter["plate"])
	}

	if filter["brand"] != "" {
		conditions += " WHERE v.brand = ? "
		args = append(args, filter["brand"])
	}

	query := `
	SELECT 
		COUNT(v.id) 
	FROM teste_gobrax.vehicle v
	LEFT JOIN teste_gobrax.driver d ON d.id = v.driver_id` + conditions

	rows, err := Db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0
	}
	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return 0
	}

	var total int
	err = rows.Scan(
		&total,
	)
	if err != nil {
		return 0
	}

	return total
}

func Create(ctx context.Context, v domain.Vehicle) (domain.Vehicle, error) {
	var err error
	var query string
	var args []interface{}
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Vehicle{}, err
	}
	if v.DriverId != 0 {
		query = `INSERT INTO vehicle (driver_id, plate, brand, model, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`
		args = append(args, v.DriverId, v.Plate, v.Brand, v.Model)
	} else {
		query = `INSERT INTO vehicle (plate, brand, model, created_at, updated_at) VALUES(?, ?, ?, NOW(), NOW())`
		args = append(args, v.Plate, v.Brand, v.Model)
	}

	result, err := Db.ExecContext(ctx, query, args...)
	if err != nil {
		return domain.Vehicle{}, err
	}
	defer Db.Close()

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	vehicle, err := GetById(ctx, int(lastId))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d", lastId)

	return vehicle, nil
}

func Update(ctx context.Context, id int, v domain.Vehicle) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `
	UPDATE vehicle
	SET 
		driver_id = ?,
		plate = ?, 
		brand = ?, 
		model = ?
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, v.DriverId, v.Plate, v.Brand, v.Model, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	return nil
}

func GetById(ctx context.Context, id int) (domain.Vehicle, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Vehicle{}, err
	}

	rows, err := Db.Query(`
	SELECT
	    id, 
		COALESCE(driver_id, 0) AS driver_id,
		plate, 
		brand, 
		model, 
		created_at, 
		updated_at  
	FROM teste_gobrax.vehicle WHERE id = ? limit 1`, id)
	if err != nil {
		return domain.Vehicle{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.Vehicle{}, errVehicleNotFound
	}
	var user domain.Vehicle
	err = rows.Scan(
		&user.Id,
		&user.DriverId,
		&user.Plate,
		&user.Brand,
		&user.Model,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return domain.Vehicle{}, err
	}

	return user, nil
}

func Delete(ctx context.Context, id int) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `DELETE FROM teste_gobrax.vehicle WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	return nil
}

func CheckPlateExist(ctx context.Context, plate string, id int) bool {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return false
	}
	var args []interface{}

	args = append(args, plate)
	conditions := ""
	if id != 0 {
		conditions += " AND id = ?"
		args = append(args, conditions)
	}

	rows, err := Db.Query(`
	SELECT 1
	FROM teste_gobrax.vehicle 
	WHERE plate = ?
	`+conditions+`
	`, args...)
	if err != nil {
		return false
	}

	defer rows.Close()

	rowExist := rows.Next()

	return rowExist
}

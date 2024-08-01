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
	errUserNotFound = errors.New("user not found")
)

func Get(ctx context.Context) ([]domain.Vehicle, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		id, 
		COALESCE(driver_id, 0), 
		placa, 
		marca, 
		modelo, 
		created_at, 
		updated_at 
	FROM teste_gobrax.vehicle`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Vehicle
	for rows.Next() {
		var user domain.Vehicle
		err := rows.Scan(
			&user.Id,
			&user.DriverId,
			&user.Placa,
			&user.Marca,
			&user.Modelo,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Create(ctx context.Context, v domain.Vehicle) (domain.Vehicle, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Vehicle{}, err
	}

	query := `INSERT INTO vehicle (placa, marca, modelo, created_at, updated_at) VALUES(?, ?, ?, NOW(), NOW())`

	result, err := Db.ExecContext(ctx, query, v.Placa, v.Marca, v.Modelo)
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
		placa = ?, 
		marca = ?, 
		modelo = ?, 
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, v.Placa, v.Marca, v.Modelo, id)
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
		placa, 
		marca, 
		modelo, 
		created_at, 
		updated_at  
	FROM teste_gobrax.vehicle WHERE id = ? limit 1`, id)
	if err != nil {
		return domain.Vehicle{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.Vehicle{}, errUserNotFound
	}
	var user domain.Vehicle
	err = rows.Scan(
		&user.Id,
		&user.DriverId,
		&user.Placa,
		&user.Marca,
		&user.Modelo,
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

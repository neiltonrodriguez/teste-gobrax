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
	errDriverNotFound = errors.New("user not found")
)

func Get(ctx context.Context) ([]domain.Driver, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		id, 
		name, 
		drivers_license, 
		phone, 
		age, 
		COALESCE('address', ''), 
		COALESCE('zipcode', ''), 
		COALESCE('neighborhood', ''), 
		COALESCE('city', ''), 
		COALESCE('uf', ''), 
		COALESCE('number', ''), 
		created_at, 
		updated_at 
	FROM driver`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Driver
	for rows.Next() {
		var user domain.Driver
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.DriversLicense,
			&user.Phone,
			&user.Age,
			&user.Address,
			&user.Zipcode,
			&user.Neighborhood,
			&user.City,
			&user.UF,
			&user.Number,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func Create(ctx context.Context, u domain.Driver) (domain.Driver, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Driver{}, err
	}

	query := `INSERT INTO driver (name, drivers_license, phone, age, zipcode, UF, address, number, neighborhood, city, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	result, err := Db.ExecContext(ctx, query, u.Name, u.DriversLicense, u.Phone, u.Age, u.Zipcode, u.UF, u.Address, u.Number, u.Neighborhood, u.City)
	if err != nil {
		return domain.Driver{}, err
	}
	defer Db.Close()

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	user, err := GetById(ctx, int(lastId))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d", lastId)

	return user, nil
}

func Update(ctx context.Context, id int, u domain.Driver) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `
	UPDATE driver 
	SET 
		name = ?, 
		drivers_license = ?, 
		phone = ?, 
		age = ?,
		address = ?,
		zipcode = ?,
		neighborhood = ?,
		city = ?,
		uf = ?,
		number = ?
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, u.Name, u.DriversLicense, u.Phone, u.Age, u.Address, u.Zipcode, u.Neighborhood, u.City, u.UF, u.Number, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	return nil
}

func GetById(ctx context.Context, id int) (domain.Driver, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Driver{}, err
	}

	rows, err := Db.Query(`
	SELECT
	    id, 
		name, 
		drivers_license, 
		phone, 
		age, 
		COALESCE('address', ''), 
		COALESCE('zipcode', ''), 
		COALESCE('neighborhood', ''), 
		COALESCE('city', ''), 
		COALESCE('uf', ''), 
		COALESCE('number', ''), 
		created_at, 
		updated_at  
	FROM teste_gobrax.driver WHERE id = ? limit 1`, id)
	if err != nil {
		return domain.Driver{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.Driver{}, errDriverNotFound
	}
	var user domain.Driver
	err = rows.Scan(
		&user.Id,
		&user.Name,
		&user.DriversLicense,
		&user.Phone,
		&user.Age,
		&user.Address,
		&user.Zipcode,
		&user.Neighborhood,
		&user.City,
		&user.UF,
		&user.Number,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return domain.Driver{}, err
	}

	return user, nil
}

func Delete(ctx context.Context, id int) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `DELETE FROM teste_gobrax.driver WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer Db.Close()

	return nil
}

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"teste-gobrax/internal/domain"
	"teste-gobrax/pkg/database"
)

var Db *sql.DB

var (
	errDriverNotFound = errors.New("driver not found")
)

func Get(ctx context.Context, filter map[string]string, pg domain.Pagination) ([]domain.Driver, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	var args []interface{}
	conditions := ""
	if filter["available"] == "true" {
		conditions += " WHERE v.id IS NULL "
	}

	var searchTerm string
	if filter["name"] != "" {
		trimKeywords := regexp.MustCompile(`[\"\*]`).ReplaceAllString(filter["name"], "")
		searchTerm = `%` + trimKeywords + `%`
		args = append(args, searchTerm)
		conditions += " WHERE d.name LIKE ? "

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
		d.id, 
		d.name, 
		d.drivers_license, 
		d.phone, 
		d.age, 
		d.created_at, 
		d.updated_at 
	FROM driver d
	LEFT JOIN teste_gobrax.vehicle v ON v.driver_id = d.id` + conditions + pagination

	rows, err := Db.QueryContext(ctx, query, args...)
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
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetTotal(ctx context.Context, filter map[string]string) int {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return 0
	}
	var args []interface{}
	conditions := ""
	if filter["available"] == "true" {
		conditions += " WHERE v.id IS NULL "
	}

	rows, err := Db.Query(`
	SELECT 
		COUNT(d.id) 
	FROM driver d
	LEFT JOIN teste_gobrax.vehicle v ON v.driver_id = d.id 
	`+conditions+` 
	`, args...,
	)
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

func CheckDriverIsAvailable(ctx context.Context, id int) bool {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return false
	}

	rows, err := Db.Query(`
	SELECT 
		d.id, 
		d.name, 
		d.drivers_license, 
		d.phone, 
		d.age, 
		d.created_at, 
		d.updated_at 
	FROM driver d
	LEFT JOIN teste_gobrax.vehicle v ON v.driver_id = d.id
	WHERE v.id IS NULL
	AND d.id = ?`, id)
	if err != nil {
		return false
	}

	defer rows.Close()

	rowExist := rows.Next()

	return rowExist
}

func Create(ctx context.Context, u domain.Driver) (domain.Driver, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Driver{}, err
	}

	query := `INSERT INTO driver (name, drivers_license, phone, age, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`

	result, err := Db.ExecContext(ctx, query, u.Name, u.DriversLicense, u.Phone, u.Age)
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
		age = ?
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, u.Name, u.DriversLicense, u.Phone, u.Age, id)
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
	var driver domain.Driver
	err = rows.Scan(
		&driver.Id,
		&driver.Name,
		&driver.DriversLicense,
		&driver.Phone,
		&driver.Age,
		&driver.CreatedAt,
		&driver.UpdatedAt)
	if err != nil {
		return domain.Driver{}, err
	}

	return driver, nil
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

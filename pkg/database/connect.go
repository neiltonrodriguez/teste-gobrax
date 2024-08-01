package database

import (
	"database/sql"
	"fmt"
	"teste-gobrax/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.GlobalConfig.Username, config.GlobalConfig.Password, config.GlobalConfig.Host, config.GlobalConfig.Port, config.GlobalConfig.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

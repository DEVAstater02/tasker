package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Username string
	Password string
	HostUrl  string
	DBPort   string
	Database string
}

func (d *DBConfig) InitDB() (*sqlx.DB, error) {
	dataSourceString := fmt.Sprintf("%s@(%s:%s)/%s", d.Username, d.HostUrl, d.DBPort, d.Database)

	dbConn, err := sqlx.Connect("mysql", dataSourceString)
	if err != nil {
		log.Fatalln(err)
	}

	return dbConn, nil
}

func (d *DBConfig) AddTask() error {
	return nil
}

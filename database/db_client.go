package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leandrorochaadm/time-register/config"
)

var DBClient *sql.DB

func InitialiseDBConnection() {
	db, err := sql.Open("mysql", config.StringConnectionDB)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}

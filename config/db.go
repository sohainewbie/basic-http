package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const PORT = "8000"

func Db() (*sql.DB, error) {
	var (
		username = "root"
		password = ""
		dbname   = "uploader"
	)
	/// Open up our database connection.
	connect := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbname)
	con, err := sql.Open("mysql", connect)

	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	return con, nil
}

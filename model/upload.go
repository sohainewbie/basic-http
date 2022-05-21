package model

import (
	con "basic-http/config"
	"fmt"
)

func InsertLogFiles(size int64, userAgent, filename, contentType string) (string, error) {
	//insert data to db
	db, err := con.Db()
	defer db.Close()
	if err != nil {
		return "", err
	}
	urlFile := "http://localhost:8000/files/" + filename
	insert, err := db.Query(fmt.Sprintf(`INSERT INTO log_files SET 
			contentType='%s', 
			size=%d,
			userAgent='%s',
			urlFile='%s'
		`, contentType, size, userAgent, urlFile))

	if err != nil {
		return "", err
	}
	defer insert.Close()
	return urlFile, nil
}

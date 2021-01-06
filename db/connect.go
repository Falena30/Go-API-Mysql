package db

import (
	"database/sql"

	//digunakan untuk menghubungkan ke mysql
	_ "github.com/go-sql-driver/mysql"
)

//Conn digunakan untuk menghubukan ke db
func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Go-API-SQL")
	if err != nil {
		return nil, err
	}

	return db, nil
}

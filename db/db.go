package db

import (
	"database/sql"
	//driver mysql
	_"github.com/go-sql-driver/mysql"
)

//ConnectDatabase tries to connect to a mysql database
func ConnectDatabase() *sql.DB{
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/things")
	if err != nil {
        panic(err.Error())
	}
    return db
}
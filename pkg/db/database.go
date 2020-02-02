package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

const DbPassword = 123
const DbUser = "root"
const DbHost = "127.0.0.1"
const DbPort = 3306

var Auth = DbUser + ":" + strconv.Itoa(DbPassword) + "@tcp(" + DbHost + ":" + strconv.Itoa(DbPort) + ")/auth"

var Con *sql.DB

func Open() {
	var err error
	Con, err = sql.Open("mysql", Auth)

	if err != nil {
		log.Fatal(err.Error())
	}
}

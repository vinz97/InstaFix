package main

import (
	"database/sql"
)

var sqlDB *sql.DB

func main() {

	//DB
	sqlDB= startDB()
	esempioQuery(sqlDB)
	// server
	startServer()


}



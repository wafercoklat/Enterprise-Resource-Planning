package datasource

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func InitDBConn() {
	str := "root:@tcp(localhost:3306)/enterprise022?parseTime=true"
	db, err := sql.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	// Check Ping for Database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Set connection to var _DBClient
	DBClient = db
}

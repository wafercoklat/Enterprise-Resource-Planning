package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitDBConn() {
	str := "root:@tcp(localhost:3306)/enterprise022?parseTime=true"
	db, err := sqlx.Open("mysql", str)
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

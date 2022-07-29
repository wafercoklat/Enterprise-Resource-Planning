package datasource

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGet_SQL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("Select * from accounts").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	// Execute
	sql := NewData()
	_, err = sql.Get("1")
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

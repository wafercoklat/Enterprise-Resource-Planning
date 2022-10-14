package database

import (
	"STACK-ERP/port"
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

// Implement
func New(dialect, dsn string, idleConn, maxConn int) (port.PortRepo, error) {
	db, err := sqlx.Open(dialect, dsn)
	if err != nil {
		fmt.Printf("Database Error - Connection - Cannot Create Connection to Database Err: %s", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Database Error - Ping Connection - Bad Ping to Database Err: %s", err)
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &Repository{db}, nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) FindByID(id string, model interface{}, tbl string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("SELECT * FROM %s WHERE ID = ?", tbl)

	err := r.db.GetContext(ctx, model, qry, id)
	// defer r.Close()

	if err != nil {
		fmt.Printf("Database Error - List - Cannot Pull Data. Table %s Err: %s", tbl, err)
		return model, err
	}

	return model, nil
}

func (r *Repository) List(data interface{}, tbl string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("SELECT * FROM %s", tbl)

	err := r.db.SelectContext(ctx, data, qry)
	if err != nil {
		fmt.Printf("Database Error - List - Cannot Pull All Data. Table %s Err: %s", tbl, err)
		return nil, err
	}

	return data, nil
}

func (r *Repository) Create(data interface{}, table, column, value string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert Data
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, column, value)
	row, err := r.db.NamedExecContext(ctx, query, data)
	if err != nil {
		fmt.Printf("Database Error - Create - Cannot Insert Data. Table %s Err: %s", table, err)
		return 0, err
	}
	return row.LastInsertId()
}

func (r *Repository) Update(data interface{}, id, table, column, columnvalue string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("UPDATE %s SET %s WHERE %s = %s", table, columnvalue, column, id)

	row, err := r.db.NamedExecContext(ctx, qry, data)
	if err != nil {
		fmt.Printf("Database Error - Update - Cannot Update Data. Table %s Err: %s", table, err)
		return 0, err
	}

	return row.RowsAffected()
}

func (r *Repository) Delete(id, tbl, column string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tbl, column)

	row, err := r.db.ExecContext(ctx, qry, id)
	if err != nil {
		fmt.Printf("Database Error - Delete - Cannot Delete Data. Table %s Err: %s", tbl, err)
		return 0, err
	}

	return row.RowsAffected()
}

func (r *Repository) Auth(uname string, model interface{}, tbl string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("SELECT UName, Password FROM %s WHERE UName = ? LIMIT 1", tbl)

	err := r.db.GetContext(ctx, model, qry, uname)
	// defer r.Close()

	if err != nil {
		return model, err
	}

	return model, nil
}

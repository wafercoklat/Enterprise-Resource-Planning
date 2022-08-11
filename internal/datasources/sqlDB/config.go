package datasources

import (
	"REVAMP-PHP-GO/internal/domain/ports"
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
func New(dialect, dsn string, idleConn, maxConn int) (ports.PortRepo, error) {
	db, err := sqlx.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
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

	qry := fmt.Sprintf("SELECT * FROM %s WHERE id = ? LIMIT 1", tbl)

	err := r.db.GetContext(ctx, model, qry, id)
	if err != nil {
		fmt.Printf("Database Error - List - Cannot Pull Data. Table %s Err: %s", tbl, err)
	}

	return model, nil
}

func (r *Repository) List(model interface{}, tbl string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("SELECT * FROM %s", tbl)

	err := r.db.SelectContext(ctx, model, qry)
	if err != nil {
		fmt.Printf("Database Error - List - Cannot Pull All Data. Table %s Err: %s", tbl, err)
	}

	return model, nil
}

func (r *Repository) Create(model interface{}, tbl, col, val string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tbl, col, val)

	row, err := r.db.NamedExecContext(ctx, qry, model)
	if err != nil {
		fmt.Printf("Database Error - Create - Cannot Insert Data. Table %s Err: %s", tbl, err)
	}

	return row.LastInsertId()
}

func (r *Repository) Update(model interface{}, id, tbl, val string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s", tbl, val, id)

	row, err := r.db.NamedExecContext(ctx, qry, model)
	if err != nil {
		fmt.Printf("Database Error - Update - Cannot Update Data. Table %s Err: %s", tbl, err)
	}

	return row.RowsAffected()
}

func (r *Repository) Delete(id, tbl string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qry := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tbl)

	row, err := r.db.ExecContext(ctx, qry, id)
	if err != nil {
		fmt.Printf("Database Error - Delete - Cannot Delete Data. Table %s Err: %s", tbl, err)
	}

	return row.RowsAffected()
}

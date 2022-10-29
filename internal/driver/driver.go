package driver

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds database connection pool.
type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

const MaxOpenConns = 10
const MaxIdleConns = 5
const MaxConnsLifetime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	if !TestConnection(db) {
		return nil, errors.New("unable to ping new database connection")
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetConnMaxLifetime(MaxConnsLifetime)
	dbconn.SQL = db

	return dbconn, nil
}

// NewDatabase creates new db connection.
func NewDatabase(dsn string) (*sql.DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
		return nil, err
	}
	if !TestConnection(d) {
		return nil, errors.New("unable to ping new database connection")
	}

	return d, err
}

// TestConnection pings a db connection to check it aliveness.
func TestConnection(db *sql.DB) bool {
	err := db.Ping()
	if err != nil {
		return false
	}
	return true
}

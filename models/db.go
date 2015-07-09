package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=lss_db user=lss_admin password=secret sslmode=disable")
	return db, err
}

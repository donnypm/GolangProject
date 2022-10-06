package config

import "database/sql"

func GetDB()(db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=localhost; user id=sa; password=Donny12345; database=test")
	return
}
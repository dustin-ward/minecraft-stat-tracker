package data

import (
	"database/sql"

	"github.com/dustin-ward/minecraft-time-logging/util"
	_ "github.com/mattn/go-sqlite3"
)

func Setup() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	util.ErrorCheck(err)
	defer db.Close()

	statement, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT,
			total_time FLOAT
		)`,
	)
	util.ErrorCheck(err)
	statement.Exec()

	return db
}

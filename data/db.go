package data

import (
	"database/sql"
	"fmt"

	"github.com/dustin-ward/minecraft-time-logging/util"
	_ "github.com/go-sql-driver/mysql"
)

// .Format("2006-01-02 15:04:05")

var db *sql.DB
var err error

func Setup() {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", username, password, host, dbname))
	util.ErrorCheck(err)
}

func Takedown() {
	db.Close()
}

func CreateUser(username string) {
	_, err := db.Query("REPLACE INTO players SET username=?", username)
	util.ErrorCheck(err)

	util.GetFace(username)
}

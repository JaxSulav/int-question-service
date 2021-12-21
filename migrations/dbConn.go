package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "interviewUser"
	dbPass := "Changeme1!"
	dbName := "interviewPortal"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

package db

import (

	_ "github.com/go-sql-driver/mysql"
	"database/sql"

	"log"
)

var db *sql.DB

func init() {
	log.Println("dsdsd");
	db, _ = sql.Open("mysql", "tangtao:123456@tcp(172.30.121.158:3307)/sampledb?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func Get() *sql.DB {

	return db;
}
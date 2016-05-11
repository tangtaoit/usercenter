package main

import (
	"net/http"
	"io"
	"fmt"
	"database/sql"
	"log"
	 _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	log.Println("dsdsd");
	db, _ = sql.Open("mysql", "tangtao:123456@tcp(172.30.121.158:3306)/sampledb?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func main() {
	//mux:=httpim.NewServeMux();


	http.HandleFunc("/user/login", userLogin)

	http.ListenAndServe(":8080", nil)
}

func CheckErr(err error)  {
	if err != nil {
		panic(err)
	}
}

func  userLogin(w http.ResponseWriter, r *http.Request){


	rows, err := db.Query("select id,rid,appid from users")

	defer rows.Close()

	CheckErr(err)

	var resultStr string
	for rows.Next() {
		var id int
		var rid string
		var appid int
		err = rows.Scan(&id,&rid,&appid)
		CheckErr(err)
		resultStr =fmt.Sprint(id,"-",rid,"-",appid)
	}

	io.WriteString(w,resultStr);

}

package main

import (
	"net/http"
	"io"
	"fmt"
	"usercenter/db"
)


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


	rows, err := db.Get().Query("select id,rid,appid from users")

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

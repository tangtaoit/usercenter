package main

import (
	"net/http"
	"io"
)

type User struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	//mux:=httpim.NewServeMux();

	http.HandleFunc("/user/login", userLogin)

	http.ListenAndServe(":8080", nil)
}

func  userLogin(w http.ResponseWriter, r *http.Request){

	io.WriteString(w,"testdd");

}

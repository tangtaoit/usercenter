package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func GetRouters() *mux.Router {

	return  NewRouter([]Route{
		Route{
			"GetUserInfo",
			"GET",
			"/users/auth",
			GetUserInfo,
		},
		Route{
			"BindUserInfo",
			"POST",
			"/users/auth",
			BindUserInfo,
		},
		Route{
			"SubmitApp",
			"POST",
			"/users/app",
			SubmitApp,
		},
	})
}

func main() {


	
	log.Fatal(http.ListenAndServe(":8080", GetRouters()))
}


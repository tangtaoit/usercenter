package main

import (
	"net/http"
	"log"
)



func main() {

	router := NewRouter([]Route{
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
	log.Fatal(http.ListenAndServe(":8080", router))
}


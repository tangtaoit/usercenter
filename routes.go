package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)


		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)

	}

	return router
}

var routes = Routes{
	Route{
		"GetUserInfo",
		"GET",
		"/users/auth/{app_id}/{app_key}",
		GetUserInfo,
	},
	Route{
		"BindUserInfo",
		"POST",
		"/users/auth/{app_id}/{app_key}",
		BindUserInfo,
	},
}
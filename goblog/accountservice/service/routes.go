package service

import "net/http"

//Route balabala
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes balabala
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount", // Name
		"GET",        //HTTP Method
		"/accounts/{accountId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}

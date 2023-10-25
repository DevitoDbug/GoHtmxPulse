package routes

import (
	"GoHtmxPulse/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, counter *int) {
	r.HandleFunc("/counter/increment",
		func(w http.ResponseWriter, r *http.Request) {
			controllers.HandleIncrement(w, r, counter)
		}).Methods("GET")

	r.HandleFunc("/counter/decrement",
		func(w http.ResponseWriter, r *http.Request) {
			controllers.HandleDecrement(w, r, counter)
		}).Methods("GET")
}

package routes

import (
	"GoHtmxPulse/controllers"
	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/increment", controllers.HandleIncrement).Methods("GET")
	r.HandleFunc("/decrement", controllers.HandleDecrement).Methods("GET")
}

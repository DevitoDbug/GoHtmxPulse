package routes

import "github.com/gorilla/mux"

func registerRoutes(r *mux.Router) {
	r.Handle("/increment", controllers.HandleIncremet()).Methods("GET")
	r.Handle("/decrement", controllers.HandleIncremet()).Methods("GET")
}

package main

import (
	"GoHtmxPulse/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := ":8000"
	counter := 0
	r := mux.NewRouter()
	routes.RegisterRoutes(r, &counter)

	r.Handle("/", r)
	r.PathPrefix("/counter/").Handler(http.StripPrefix("/counter/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server at port ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}

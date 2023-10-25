package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := ":8080"
	r := mux.NewRouter()

	r.Handle("/", r)

	fmt.Println("Starting server at port ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}

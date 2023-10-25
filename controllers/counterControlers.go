package controllers

import "net/http"

func HandleIncrement(w http.ResponseWriter, r *http.Request, counter *int) {
	*counter++
	w.WriteHeader(http.StatusOK)
}

func HandleDecrement(w http.ResponseWriter, r *http.Request, counter *int) {
	*counter--
	w.WriteHeader(http.StatusOK)
}

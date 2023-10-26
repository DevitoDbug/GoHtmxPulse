package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandleIncrement(w http.ResponseWriter, r *http.Request, counter *int) {
	*counter++
	temp := template.Must(template.New("counter").Parse(fmt.Sprintf("<div id = \"counter\" class=\"\">%d</div>\n", *counter)))
	temp.ExecuteTemplate(w, "counter", *counter)
	w.Header().Set("Content-Type", "application/json")
}

func HandleDecrement(w http.ResponseWriter, r *http.Request, counter *int) {
	*counter--
	temp := template.Must(template.New("counter").Parse(fmt.Sprintf("<div id = \"counter\" class=\"\">%d</div>\n", *counter)))
	temp.ExecuteTemplate(w, "counter", nil)
	w.Header().Set("Content-Type", "application/json")
}

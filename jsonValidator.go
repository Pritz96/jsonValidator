package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		r.ParseForm()
		jsonInput := r.FormValue("json")
		outcome := json.Valid([]byte(jsonInput))
		if outcome {
			fmt.Fprintf(w, "Valid JSON")
		} else {
			fmt.Fprintf(w, "Invalid JSON")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(indexHandler))
	mux.Handle("/formHandler", http.HandlerFunc(formHandler))
	log.Println("Starting JSON Validator on :8080 ...")
	http.ListenAndServe(":8080", mux)
}

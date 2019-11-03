package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	mux := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles("form.html"))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}))
	mux.Handle("/formHandler", http.HandlerFunc(formHandler))
	http.ListenAndServe(":8080", mux)
}

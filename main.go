package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world...")

	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

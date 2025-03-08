package main

import (
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		ctx := map[string][]Article{
			"Articles": {
				{Title: "First article", Content: "hello world!"},
				{Title: "Second article", Content: "hello world, again!"},
			},
		}

		tmpl.Execute(w, ctx)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

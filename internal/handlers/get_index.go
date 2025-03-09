package handlers

import (
	"html/template"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func Get_Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	ctx := map[string][]Article{
		"Articles": {
			{Title: "First article", Content: "hello world!"},
			{Title: "Second article", Content: "hello world, again!"},
		},
	}

	tmpl.ExecuteTemplate(w, "homepage", ctx)
}

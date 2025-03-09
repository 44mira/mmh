package main

import (
	"github.com/44mira/mmh/internal/handlers"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/", handlers.Get_Index)
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

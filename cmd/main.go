package main

import (
	"github.com/44mira/mmh/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Get_Index)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

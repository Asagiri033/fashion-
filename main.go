package main

import (
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"registration/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

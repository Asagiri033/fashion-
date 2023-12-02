package main

import (
	"database/sql"
	"log"
	"net/http"

	"fashion/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	dataSourceName := "root:root@tcp(localhost:3306)/fashion_db"
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/users", services.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", services.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users/{id}", services.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", services.DeleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", helloMux).Methods("GET")
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	r.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


func main() {
	fmt.Println("Gorm revisit")

	InitialMigration()

	handleRequest()
}

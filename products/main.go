package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initDB()

	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", createOrder).Methods("POST")
	// Read
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	// Read-all
	router.HandleFunc("/orders", getOrders).Methods("GET")
	// Update
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
	// Initialize db connection


	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {

	http.HandleFunc("/", helloworld)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

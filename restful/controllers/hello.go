package controllers

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

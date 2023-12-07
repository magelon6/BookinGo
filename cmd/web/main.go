package main

import (
	"fmt"
	"net/http"

	"github.com/magelon6/bookinGo/pkg/handlers"
)

const portNumber = ":8080"

// Main app func
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
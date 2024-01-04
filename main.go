package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	// start the server on port 8000'``
	fmt.Println("Server is running at 8000 port")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
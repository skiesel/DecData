package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Access Server")
	http.HandleFunc("/register", register)
	http.HandleFunc("/syncAccess", syncAccess)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

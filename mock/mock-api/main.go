package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /rest/api/latest/encrypt", authenticationMiddleware(encryptString))
	mux.HandleFunc("POST /api/passwords", storePassword)

	log.Println("Listening on 6969")
	http.ListenAndServe(":6969", mux)

}

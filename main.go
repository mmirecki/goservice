package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {

	log.Print("Starting the service...")

	id := rand.Int()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Port is not set. Using 8000")
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed. HOME  /  ", id)
	},
	)

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello1! Your request was processed.   /   ", id)
	},
	)

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

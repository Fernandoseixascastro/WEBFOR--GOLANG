package main

import (
	"PageWeb/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}

}

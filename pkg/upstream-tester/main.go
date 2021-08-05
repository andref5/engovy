package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = "5000"

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("PONG")
		fmt.Fprintf(w, "PONG")
	})
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		log.Println("PING")
		fmt.Fprintf(w, "PING")
	})

	log.Println("Upstream tester server startup port " + Port)
	http.ListenAndServe(":"+Port, nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("|  %s  |  %s  |  %s", r.Host, r.Method, r.RequestURI)
	_, err := fmt.Fprintf(w, "Message: %v", r.URL.Path[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("Start service")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 8080, "The port that the server listens to")
)

func main() {
	flag.Parse()
	log.Print("Starting HTTP server on port ", *port)

	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

	if err != nil {
		log.Fatal("HTTP server error: ", err)
	}
}

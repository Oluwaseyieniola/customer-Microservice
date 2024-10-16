package main

import (
	"log"
	"net/http"
	 

	"github.com/Oluwaseyieniola/mgdonalds/commons"
)

var (
	httpAddr = commons.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := Newhandler()
	handler.registerRoutes(mux)

	log.Printf("Starting the server at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server...")
	}
}

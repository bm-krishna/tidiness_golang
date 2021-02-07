package server

import (
	"log"
	"net/http"
	"time"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/handlers"
)

func BootStrapServer() {
	// create handler
	handler := &handlers.HandlerService{}
	// set up server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server running on port", server.Addr)
	log.Fatal(server.ListenAndServe())
}

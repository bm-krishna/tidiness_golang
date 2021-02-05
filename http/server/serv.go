package server

import (
	"log"
	"net/http"
	"os"

	"github.com/bm-krishna/tidiness_golang/http/handlers"
)

func Serve() {
	log.Println("Main Execution")
	logger := log.New(os.Stdout, "login-handler", log.LstdFlags)
	loginHandler := handlers.LoginH(logger)

	serveMux := http.NewServeMux()

	serveMux.Handle("/login", loginHandler)
	server := &http.Server{}

	http.ListenAndServe(":9090", serveMux)

}

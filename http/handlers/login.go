package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	logger *log.Logger
}

func LoginH(logger *log.Logger) *Login {
	return &Login{
		logger,
	}
}

func (login *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	login.logger.Println("Login handler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Failed to Read body from Request", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "%s", body)
}

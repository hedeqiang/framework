package main

import (
	"github.com/hedeqiang/framework/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8081",
	}
	server.ListenAndServe()
}

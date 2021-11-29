package main

import (
	"net/http"
	"github.com/hedeqiang/hike/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8081",
	}
	server.ListenAndServe()
}

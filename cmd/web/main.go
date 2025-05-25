package main

import (
	"log"
	"net/http"
	"time"
)

type application struct{}

const port = ":4000"

func main() {
	app := application{}
	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       time.Second * 30,
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
	}

	log.Printf("Server is running on port %v", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

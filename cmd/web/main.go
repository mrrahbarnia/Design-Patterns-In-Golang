package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/configuration"
)

type application struct {
	config appConfig
	App    *configuration.Application
}
type appConfig struct {
	dsn string
}

const port = ":4000"

func main() {
	cfg := &appConfig{
		dsn: "postgres://admin:123456@localhost:5432/db?sslmode=disable",
	}
	db, err := initPostgresDB(cfg.dsn)
	if err != nil {
		log.Panic(err)
	}
	app := application{
		config: *cfg,
		App:    configuration.New(db),
	}

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

package main

import (
	"log"
	"net/http"
)

type application struct{}

const port = ":4000"

func main() {
	// app := application{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world."))
	})

	log.Printf("Server is running on port %v", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}

}

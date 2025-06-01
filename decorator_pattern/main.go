package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// Decorator pattern for logging time taken by the handler
func TimingDecorator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r) // Call the next handler in the chain
		log.Println(time.Since(now), "for request: ", r.URL.Path)
	})
}

func main() {
	http.Handle("/hello", TimingDecorator(http.HandlerFunc(HelloHandler)))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

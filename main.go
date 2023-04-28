package main

import (
	"log"
	"net/http"
	"time"
)

type healthHandler struct {
}

func (hh healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("te2222"))

}

func main() {
	m := http.NewServeMux()
	m.Handle("/health", healthHandler{})

	s := &http.Server{
		Handler:      m,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port 8080")
	log.Fatal(s.ListenAndServe())

}

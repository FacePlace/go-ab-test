package main

import (
	"log"
	"net/http"
	"time"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type AbHandler struct {
}

func (h *AbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	designA := []byte("<html><head><title>The Golang Hotel</title></head><body><p>The Golang Hotel is a relaxing place !</p><p>We offer 20% discount if you call this number : <strong>1234567891</strong></p></body></html>")

	designB := []byte("<html><head><title>The Golang Hotel</title></head><body><h2>The Golang Hotel is a relaxing place !</h2><h5>We offer 20% discount if you call this number : <strong>1234567892</strong></h5></body></html>")

	minutes := time.Now().Minute()

	if minutes%2 == 0 {
		log.Println("Serving A")
		_, err := w.Write(designA)
		if err != nil {
			log.Println("Impossible to serve design A", err)
		}
	} else {
		log.Println("Serving B")
		_, err := w.Write(designB)
		if err != nil {
			log.Println("Impossible to serve design B", err)
		}
	}
}

func main() {
	abTestingServer := &http.Server{
		Addr:         "127.0.0.1:9899",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &AbHandler{},
	}
	log.Fatal(abTestingServer.ListenAndServe())
}

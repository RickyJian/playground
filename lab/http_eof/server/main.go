package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	fmt.Fprintf(w, "hello world!!")
}

func main() {
	http.HandleFunc("/", index)
	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

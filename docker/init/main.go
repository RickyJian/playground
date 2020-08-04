package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my playground")
	})
	if err :=http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

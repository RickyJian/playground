package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("build request error: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}
	defer resp.Body.Close()

	var bs bytes.Buffer
	io.Copy(&bs, resp.Body)
	fmt.Println("result resp:", bs.String())
}

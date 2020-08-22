package main

import (
	"log"
	"os"
)

const (
	src  = "test/a"
	dest = "test/b"
)

func main() {
	if _, err := os.Stat(src); err != nil {
		log.Fatalf("src error: %v", err)
	} else if err := os.Rename(src, dest); err != nil {
		log.Fatalf("move file error: %v", err)
	}
}

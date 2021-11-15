package main

import (
	"log"
	"os"
)

func StartLogger() *os.File {
	f, err := os.OpenFile("numbers.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
	return f
}

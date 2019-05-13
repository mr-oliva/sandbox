package main

import (
	"log"

	"github.com/bookun/sandbox/go/error-handling-iota/sample"
)

func main() {
	if errCode, err := sample.Sample(3); err != nil {
		log.Fatalf("errCode: %v, err: %v\n", errCode, err)
	}
}

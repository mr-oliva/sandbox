package main

import (
	"net/http"

	"github.com/bookun/sandbox/go/reply-body/function"
)

func main() {
	http.HandleFunc("/", function.GetBody)
	http.ListenAndServe(":8080", nil)
}

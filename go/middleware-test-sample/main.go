package main

import (
	"fmt"
	"net/http"

	"github.com/bookun/sandbox/go/middleware-test-sample/middleware"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main handler")
}

func main() {
	http.HandleFunc("/", middleware.M1(hoge))
	http.ListenAndServe(":8080", nil)
}

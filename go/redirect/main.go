package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirct(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "https://www.google.com/", 302)
	http.Redirect(w, r, "/second", 302)
}

func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "second")
}

func main() {
	http.HandleFunc("/", redirct)
	http.HandleFunc("/second", second)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

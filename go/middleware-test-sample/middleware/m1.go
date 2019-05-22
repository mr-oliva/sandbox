package middleware

import (
	"fmt"
	"net/http"
)

func M1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "m1 start")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "m1 end")
	}
}

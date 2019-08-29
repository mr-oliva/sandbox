package middleware

import (
	"context"
	"net/http"
)

func M1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context.WithValue(r.Context(), "hoge", "hoge")
	}
}

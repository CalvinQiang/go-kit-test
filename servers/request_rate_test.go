package main

import (
	"golang.org/x/time/rate"
	"net/http"
	"testing"
)

var r = rate.NewLimiter(1, 6)

func MyLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !r.Allow() {
			http.Error(writer, "to many requests", http.StatusTooManyRequests)
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
func TestRequestRate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	http.ListenAndServe(":8111", MyLimit(mux))
}

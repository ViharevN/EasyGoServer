package main

import (
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandle)))

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}

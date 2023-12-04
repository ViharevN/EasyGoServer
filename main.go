package main

import "net/http"

func HandleFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "ok"}`))
}

func main() {
	http.HandleFunc("status", HandleFunc)
	http.ListenAndServe(`:8080`, nil)
}

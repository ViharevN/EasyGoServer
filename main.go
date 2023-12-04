package main

import "net/http"

func StatusHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "ok"`))
}

func main() {
	http.HandleFunc("/status", StatusHandler)
	http.ListenAndServe(`:8080`, nil)
}

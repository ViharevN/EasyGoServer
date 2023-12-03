package main

import "net/http"

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ya.ru", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/search", redirect)
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}

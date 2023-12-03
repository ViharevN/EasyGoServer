package main

import "net/http"

func handle(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./cmd/main.go")
	})
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}

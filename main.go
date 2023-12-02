package main

import "net/http"

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello")
	res.Write(data)
}

func main() {
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}

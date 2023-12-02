package main

import "net/http"

func mainPage(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Hello!"))
}

func apiPage(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("This is api page/"))
}

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/api", apiPage)

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}

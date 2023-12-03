package main

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ============\r\n"
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	res.Write([]byte(body))
}

func apiPage(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("This is api page/"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api/", apiPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", http.NoBody)
	if err != nil {
		panic(err)
	}
	request.Header.Set(`MyHeader`, "Hello")
	request.Header.Add(`MyHeader`, "Привет")

	response, err := client.Do(request)
	io.Copy(os.Stdout, response.Body)
	defer response.Body.Close()

}

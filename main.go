package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	var body = []byte(`{"message":"Hello"}`)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}

package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	var fileName = "./Read.ME"
	client := &http.Client{}

	file, _ := os.Open(fileName)

	defer file.Close()
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("uploadfile", fileName)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	writer.Close()

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", body)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := client.Do(request)
	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)

	fmt.Println(string(res))
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 512 && i < len(body); i++ {
		fmt.Printf("%d ", body[i])
	}

	fmt.Printf("\r\nStatus code: %d\r\n", response.StatusCode)
	for k, v := range response.Header {
		fmt.Printf("%s: %v\r\n", k, v[0])
	}
}

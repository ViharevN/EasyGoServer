package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	defer response.Body.Close()
	if _, err = io.ReadAll(response.Body); err != nil {
		fmt.Println(err)
	}
	for s, v := range response.Header {
		fmt.Printf("%s: %v\r\n", s, v)
	}

}

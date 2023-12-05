package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(1),
	}

	data := url.Values{}
	data.Set("key1", "value1")
	data.Set("key2", "value2")

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/x-www-urlencoded")
	request.Header.Set("Content-Type", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(request)

	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

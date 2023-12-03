package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	subj := Subj{"Milk", 50}

	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин <input type="text" name="login"></label>
            <label>Пароль <input type="password" name="password"></label>
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Welcome")
		} else {
			http.Error(w, "incorrect login or password", http.StatusUnauthorized)
		}
	} else {
		io.WriteString(w, form)
	}
}

func main() {
	err := http.ListenAndServe(`:8080`, http.HandlerFunc(mainPage))
	if err != nil {
		panic(err)
	}
}

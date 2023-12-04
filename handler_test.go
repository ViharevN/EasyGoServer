package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive Test #1",
			want: want{
				code:        200,
				response:    `{"status": "ok"`,
				contentType: "application/json",
			},
		}, {
			name: "negative Test #2",
			want: want{
				code:        404,
				response:    "Not Found",
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/status", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			StatusHandler(rr, req)

			if status := rr.Code; status != test.want.code {
				t.Errorf("Ожидался код ответа %v, получено %v", test.want.code, status)
			}

			if body := rr.Body.String(); body != test.want.response {
				t.Errorf("Ожидалось тело ответа %v, получено %v", test.want.response, body)
			}

			if ct := rr.Header().Get("Content-Type"); ct != test.want.contentType {
				t.Errorf("Ожидался Content-Type %v, получено %v", test.want.contentType, ct)
			}
		})

	}
}

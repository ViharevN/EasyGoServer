package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		code        int
		contentType string
		user        User
	}

	tests := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{
			name: "positive test #1",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				code:        200,
				contentType: "application/json",
				user: User{ID: "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		}, {
			name: "positive test #2",
			users: map[string]User{
				"id2": {
					ID:        "id2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			want: want{
				code:        200,
				contentType: "application/json",
				user: User{ID: "id2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(test.users))
			h(w, request)

			result := w.Result()

			assert.Equal(t, test.want.code, result.StatusCode)
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			assert.Nil(t, err)
			require.NoError(t, err)
			err = result.Body.Close()

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, test.want.user, user)
		})
	}
}

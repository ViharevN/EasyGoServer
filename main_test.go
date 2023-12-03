package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		fullName  string
	}{
		{
			name:      "testcase #1",
			firstName: "John",
			lastName:  "Johnson",
			fullName:  "John Johnson",
		}, {
			name:      "testcase #2",
			firstName: "Sergey",
			lastName:  "Sergeev",
			fullName:  "Sergey Sergeev",
		}, {
			name:      "testcase #3",
			firstName: "Georgy",
			lastName:  "Georgiev",
			fullName:  "Georgy Georgiev",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user := User{
				FirstName: test.firstName,
				LastName:  test.lastName,
			}

			v := user.FullName()
			assert.Equal(t, test.fullName, v)
		})
	}
}

package main

import "testing"

func TestFamily_AddNew(t *testing.T) {
	tests := []struct {
		name         string
		relationship Relationship
		person       Person
		expectedErr  error
	}{
		{
			name:         "testcase #1",
			relationship: Father,
			person: Person{
				FirstName: "Misha",
				LastName:  "Popov",
				Age:       56,
			},
			expectedErr: nil,
		}, {
			name:         "testcase #2",
			relationship: Father,
			person: Person{
				FirstName: "Grug",
				LastName:  "Mishi",
				Age:       57,
			},
			expectedErr: ErrRelationShipAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := Family{}
			err := f.AddNew(test.relationship, test.person)
			if err != test.expectedErr {
				t.Errorf("Expected error %v, but got %v", test.expectedErr, err)
			}
		})
	}
}

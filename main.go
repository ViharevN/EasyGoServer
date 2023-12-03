package main

import (
	"errors"
	"fmt"
)

type Relationship string

const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

type Family struct {
	Members map[Relationship]Person
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var (
	ErrRelationShipAlreadyExists = errors.New("relationship allreqdy exists")
)

func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationShipAlreadyExists
	}
	f.Members[r] = p
	return nil
}

func main() {
	f := Family{}
	err := f.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56,
	})
	fmt.Println(f, err)

	err = f.AddNew(Father, Person{
		FirstName: "Grug",
		LastName:  "Mishi",
		Age:       57,
	})
	fmt.Println(f, err)
}

package main

import (
	"fmt"
	"log"
)

// Write a program that validates cats!
// PS: Only valid cats are jessie, of course.

// Resources:
// https://go.googlesource.com/proposal/+/master/design/go2draft.md
// There are two sections - error handling and error values. Both are interesting.

type Cat struct {
	Name string
}

func validName(name string) error {
	// I am valid if i am Jessie teehee
	if name != "Jessie" {
		return fmt.Errorf("name is not jessie")
	}

	return nil
}

func NewCat(name string) (*Cat, error) {
	// uses valid name
	if err := validName(name); err != nil {
		return nil, err
	}
	return &Cat{Name: name}, nil
}

func main() {
	_, err := NewCat("Jessie2")
	log.Printf("ERr err: %v / %T", err, err)

	//log.Printf("Cat with name: %s, err: %v", cat.Name, err)
}

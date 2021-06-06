package main

import (
	"testing"
)

// Resources:
// - https://pkg.go.dev/github.com/stretchr/testify
// - https://github.com/golang/mock
// - Packages with tests:
// - https://github.com/restic/restic
// - Example of flat: https://github.com/olivere/elastic

func TestSum(t *testing.T) {
	if Sum(2, 3) != 4 {
		t.Fatal("Maths says 2 + 3 is 5")
	}
}

package main

import "fmt"

// []<type> are 0 indexed slices(expandable)
// Suggestion:
// https://gobyexample.com/slices
// https://tour.golang.org/moretypes/7

func main() {

	// For is the only loop we have and we love it.
	// Suggestion: Transfor the slice into a map and loop over it.
	for index, animal := range []string{"cat", "dog", "fawn"} {
		fmt.Printf("At position %d we have the %s\n", index, animal)
		if animal == "cat" {
			fmt.Println("...which is clearly the best animal, btw!")
		}

		switch animal {
		case "dog":
			fmt.Println("...a distant second")
		case "fawn":
			fmt.Println("...being just a silly bird")
		}
	}

}

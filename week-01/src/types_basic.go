package main

import "fmt"

const (
	// Constantas are the only immutable data we have
	BestCatName = "Jessie"
)

// Type size will be architecture dependent unless explicitely specified.
func sum(a, b int64) int64 {
	return a + b
}

// Golang pass by value, but the value of pointers allow side effects
func swapMe(s *string) {
	*s = "Bienvenado"
}

func main() {

	welcomeInFrench := "Je suis Golang! "

	fmt.Printf("%d\n", sum(5, 3))

	fmt.Printf("%t\n", true && false)

	fmt.Printf("%f\n", float64(3.0)+float64(4.2))

	fmt.Printf("%s\n", welcomeInFrench)

	fmt.Printf("I am %s the cat\n", BestCatName)

	var welcome string
	welcome = "...and I am a cat"

	fmt.Printf("%s\n", welcome)
	alsoWelcome := &welcome

	fmt.Printf("%s\n", *alsoWelcome)

	welcome = "...while I am a dog"

	swapMe(&welcome)

	fmt.Printf("%s\n", *alsoWelcome)
}

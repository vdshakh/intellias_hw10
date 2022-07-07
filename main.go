package main

import (
	"fmt"
)

func countFeed(a []Animal) int {
	var feedForAnimal int

	for _, animal := range a {
		fmt.Printf("%s kg needs %v kg of feed\n", animal.String(), animal.eatPerMonth())
		feedForAnimal += animal.eatPerMonth()
	}

	return feedForAnimal
}

func main() {
	animals := []Animal{
		Dog{weight: 12},
		Cat{weight: 7},
		Cow{weight: 88},
		Dog{weight: 23},
	}

	feedPerMonth := countFeed(animals)
	fmt.Printf("\nKG OF FEED PER MONTH IS NEEDED: %v \n", feedPerMonth)
}

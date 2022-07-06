package main

import (
	"fmt"
)

func main() {
	animals := []Animal{
		Dog{weight: 12},
		Cat{weight: 7},
		Cow{weight: 88},
		Dog{weight: 23},
	}

	for _, animal := range animals {
		fmt.Printf("%s kg needs %v kg of feed\n", animal.String(), animal.eatPerMonth())
	}

	pets := Farm{
		cattle: animals,
	}

	var feedPerMonth int

	for _, pet := range pets.cattle {
		feedPerMonth += pet.eatPerMonth()
	}

	fmt.Printf("\nKG OF FEED PER MONTH IS NEEDED: %v \n", feedPerMonth)
}

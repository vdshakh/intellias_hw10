package main

import (
	"fmt"
)

func countFeed(a []Animal) int {
	var feedForAnimal int

	for _, animal := range a {
		fmt.Printf("%s weighing %v kg needs %v kg of feed\n", animal, animal.getWeight(), animal.eatPerMonth())
		feedForAnimal += animal.eatPerMonth()
	}

	return feedForAnimal
}

func validator(a Animal) error {
	err := validateType(a) // or if err := validateType(a); err != nil {
	if err != nil {
		return fmt.Errorf("validateType failed: %v", err)
	}

	err = validateWeight(a)
	if err != nil {
		return fmt.Errorf("validateWeight failed: %v", err)
	}

	err = validateIsEatable(a)
	if err != nil {
		return fmt.Errorf("validateIsEatable failed: %v", err)
	}

	return nil
}

func validateType(a Animal) error {
	switch animalType := a.(type) {
	case Dog:
		if "Dog" != animalType.String() {
			return fmt.Errorf("typeAnimal doesn't match the animal name: %v", animalType.String())
		}
	case Cat:
		if "Cat" != animalType.String() {
			return fmt.Errorf("typeAnimal doesn't match the animal name: %v", animalType.String())
		}
	case Cow:
		if "Cow" != animalType.String() {
			return fmt.Errorf("typeAnimal doesn't match the animal name: %v", animalType.String())
		}
	default:
		return fmt.Errorf("unknown animalType: %v", animalType.String())
	}

	return nil
}

func validateWeight(a Animal) error {
	minWeight := 5
	weight := a.getWeight()

	if weight < minWeight {
		return fmt.Errorf("weight of %v is incorrect: %v", a, weight)
	}

	return nil
}

func validateIsEatable(a Animal) error {
	if _, ok := isEatableList[a.String()]; !ok {
		return fmt.Errorf("%v isn't eatable", a)
	}

	return nil
}

func main() {
	animals := []Animal{
		Dog{weight: 12},
		Cat{weight: 3},
		Cow{weight: 88},
		Dog{weight: 23},
	}

	feedPerMonth := countFeed(animals)
	fmt.Printf("\nKG OF FEED PER MONTH IS NEEDED: %v \n", feedPerMonth)

	for i, _ := range animals {
		err := validator(animals[i])
		if err != nil {
			fmt.Printf("\n\nvalidator failed: %v", err)
		}
	}
}

package main

import (
	"fmt"
	"strconv"
)

const (
	feedDog = 2
	feedCat = 7
	feedCow = 25
)

type Dog struct {
	weight int
}

func (d Dog) eatPerMonth() int {
	return d.weight * feedDog
}

func (d Dog) String() string {
	return "Dog weighing " + strconv.Itoa(d.weight)
}

type Cat struct {
	weight int
}

func (c Cat) eatPerMonth() int {
	return c.weight * feedCat
}

func (c Cat) String() string {
	return "Cat weighing " + strconv.Itoa(c.weight)
}

type Cow struct {
	weight int
}

func (cw Cow) eatPerMonth() int {
	return cw.weight * feedCow
}

func (cw Cow) String() string {
	return "Cow weighing " + strconv.Itoa(cw.weight)
}

type Animal interface {
	eatPerMonth() int
	fmt.Stringer
}

type Farm struct {
	cattle []Animal
}

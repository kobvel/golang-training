package main

import "fmt"

type Animal struct {
	Health int
	Name   string
}

func (a Animal) move() {
	fmt.Println(a.Name, " is under the way")
}

type Frog struct {
	Animal
}

func (f Frog) say() {
	fmt.Println("Quack")
}

func main() {
	frog := Frog{
		Animal: Animal{
			Health: 10,
			Name:   "Fred",
		},
	}

	frog.say()
	frog.move()
}

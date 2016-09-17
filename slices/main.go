package main

import "fmt"

func main() {
	fmt.Println("_____ LITERAL Declaration _____")

	student := []string{}
	students := [][]string{}

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(students == nil)

	fmt.Println("_____ VAR Declaration _____")

	var deal []string

	fmt.Println(deal)
	deal = append(deal, "s")
	fmt.Println(deal)

	fmt.Println("_____ MAKE Declaration _____")

	slc := make([]byte, 10, 100)
	fmt.Println(slc)
}

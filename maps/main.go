package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Buenas Diaz",
		2: "Cheers",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}

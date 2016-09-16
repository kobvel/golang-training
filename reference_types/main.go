package main

import "fmt"

func main() {
	// m := make([]string, 1, 25)
	d := []string{"het", "he"}
	changeMe(d)
	fmt.Println(d)
}

func changeMe(z []string) {
	z[1] = "John"
	fmt.Println(z)
}

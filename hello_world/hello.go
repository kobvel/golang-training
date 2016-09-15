package main

import "fmt"

const a string = "hey"

func main() {
	// fmt.Println("Hello World")

	var a int8
	var b int8
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("Result ", a/b)
}

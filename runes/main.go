package main

import "fmt"

func main() {
	for i := 5000; i < 5100; i++ {
		fmt.Printf("%v - %v \n", string(i), []byte(string(i)))
	}
}

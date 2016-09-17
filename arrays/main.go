package main

import "fmt"

func main() {
	var x [256]byte

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < len(x); i++ {
		x[i] = byte(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}

package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(n int) (int, bool) {
	r := n / 2
	if n%2 == 0 {
		return r, true
	}
	return r, false
}

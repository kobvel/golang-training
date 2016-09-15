package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}

	fmt.Println()

	i := 0
	for i < 10 {
		i++
		fmt.Print(i)
	}

	fmt.Println()

	j := 0
	for {
		j++
		if j > 50 {
			break
		}
		fmt.Print(j)
	}
}

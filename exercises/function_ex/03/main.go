package main

import "fmt"

func main() {
	fmt.Println(greatest([]int{1, 23, 4, 5, 6, 21, 4}...))
}

func greatest(nums ...int) int {
	var r int
	for _, n := range nums {
		if n > r {
			r = n
		}
	}

	return r
}

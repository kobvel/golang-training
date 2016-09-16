package main

import "fmt"

func main() {
	slc := []int{12, 15, 12, 10, 12, 12, 5}
	nc := filter(slc, func(n int) bool {
		return n != 12
	})
	fmt.Println(nc)

}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

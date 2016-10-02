package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {

	studyGroup := people{"Zeno", "Abba", "Den", "Yasha"}
	_ = sort.Reverse(studyGroup)
	fmt.Println(studyGroup)

	s := []string{"Zeno", "Abba", "Den", "Yasha"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	n := []int{7, 6, 4, 2, 1, 6, 2}
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
}

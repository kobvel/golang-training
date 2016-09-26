package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First  string
	Second string
	Age    int `json:"wisdome score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.Age)
	fmt.Println(p1.Second)
	fmt.Println(p1.First)

	bs := []byte(`{"First":"James","Second":"Bond","wisdome score":12}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1.Age)
	fmt.Println(p1.Second)
	fmt.Println(p1.First)
}

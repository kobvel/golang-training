package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First         string
	Second        string
	Age           int `json:"wisdom score"`
	IgonoredByTag int `json:"-"`
	notExported   int // will not be presented after Marshal
}

func main() {
	p1 := Person{"James", "Bond", 22, 44, 12}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}

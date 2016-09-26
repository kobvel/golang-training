package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First  string
	Second string
	Age    int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"James", "Second":"Bond", "Age":12}`)
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1)
}

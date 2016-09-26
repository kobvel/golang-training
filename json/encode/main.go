package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First  string
	Second string
	Age    int
}

func main() {
	p1 := Person{"James", "Bond", 12}
	json.NewEncoder(os.Stdout).Encode(&p1)

}

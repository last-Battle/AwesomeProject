package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	p1 := Person{
		Name:   "martin",
		Age:    18,
		Weight: 70,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		return
	}
	fmt.Printf("str: %s\n", b)

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

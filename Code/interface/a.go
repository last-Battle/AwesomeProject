package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Print() {
	fmt.Printf("Point{X=%d, Y=%d}\n", p.X, p.Y)
}

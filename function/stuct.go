package function

import "fmt"

type Circle struct {
	radius float64
}

func Struc() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("sqare area is", c1.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

package function

import (
	"fmt"
	"math"
)

func Func1() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
}

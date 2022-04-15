package function

import "fmt"

func Switch() {
	var grade string = "b"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	default:
		grade = "D"
	}
	fmt.Println("your mark is", grade)
}

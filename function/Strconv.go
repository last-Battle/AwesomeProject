package function

import (
	"fmt"
	"strconv"
)

func STRCONV() {
	s1 := "100"
	//i2 := 200
	//s2 := strconv.Itoa(i2)
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("cant conv")
	} else {
		fmt.Printf("type: %T value: %#v\n", i1, i1)
	}
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	q := strconv.Quote("Hello, 世界")
	//q := strconv.QuoteToASCII("Hello, 世界")
	println(b, f, i, u, q)
}

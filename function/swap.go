package function

//func Swap(x, y string) (string, string) {
//	return y, x
//}

func Swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

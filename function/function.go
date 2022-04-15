package function

func Max(Num1, Num2 int) int {
	var Result int
	if Num1 > Num2 {
		Result = Num1
	} else {
		Result = Num2
	}
	return Result
}

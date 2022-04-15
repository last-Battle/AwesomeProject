package function

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SCAN() {
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scan(&name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

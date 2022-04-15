package function

import (
	"fmt"
	"io"
	"os"
)

func FMT() {
	fmt.Print("Print")

	_, err := fmt.Fprintln(os.Stdout, "pls input")
	if err != nil {
		return
	}
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	name := "martin"
	_, err = fmt.Fprintf(fileObj, "input file %s", name)
	if err != nil {
		return
	}

	age := 14
	s := fmt.Sprintf("name:%s, age:%d", name, age)
	_, err = io.WriteString(os.Stdout, s)
	if err != nil {
		return
	}
	aaa := fmt.Errorf("this is an errr")
	print(aaa)
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%T	打印值的类型
	//%%	百分号
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	//%t	true或false
	//%b	表示为二进制
	//%c	该值对应的unicode码值
	//%d	表示为十进制
	//%o	表示为八进制
	//%x	表示为十六进制，使用a-f
	//%X	表示为十六进制，使用A-F
	//%U	表示为Unicode格式：U+1234，等价于”U+%04X”
	//%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	//%b	无小数部分、二进制指数的科学计数法，如-123456p-78
	//%e	科学计数法，如-1234.456e+78
	//%E	科学计数法，如-1234.456E+78
	//%f	有小数部分但无指数部分，如123.456
	//%F	等价于%f
	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	//%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	//%s	直接输出字符串或者[]byte
	//%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	//%x	每个字节用两字符十六进制数表示（使用a-f)
	//%X	每个字节用两字符十六进制数表示（使用A-F）

	//%f	默认宽度，默认精度
	//%9f	宽度9，默认精度
	//%.2f	默认宽度，精度2
	//%9.2f	宽度9，精度2
	//%9.f	宽度9，精度0
}

package function

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	//file, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Println("open file failed")
	//}
	//defer file.Close()
	//var tmp = make([]byte, 128)
	//n, err := file.Read(tmp)
	//if err == io.EOF {
	//	fmt.Println("end of the line")
	//	return
	//}
	//if err != nil {
	//	fmt.Println("read file failed, err:", err)
	//	return
	//}
	//fmt.Println("read %d bytes data\n", n)
	//fmt.Println(string(tmp[:n]))
	//file, err := os.Open("./xx.txt")
	//if err != nil {
	//	fmt.Println("open file failed, err:", err)
	//	return
	//}
	//defer file.Close()
	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		if len(line) != 0 {
	//			fmt.Println(line)
	//		}
	//		fmt.Println("文件读完了")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file failed, err:", err)
	//		return
	//	}
	//	fmt.Print(line)
	//}
	//content, err := ioutil.ReadFile("./main.go")
	//if err != nil {
	//	fmt.Println("read file failed, err", err)
	//	return
	//}
	//fmt.Println(string(content))
	str := "hell world"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

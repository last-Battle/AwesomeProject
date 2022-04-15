package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const txtPath = "/txt/"

type fileHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(h fileHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func GetFile(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len(txtPath):]
	file, err := os.Open(path)
	if err != nil {
		// fmt.Fprintf(w, "%s\n", err.Error())
		return err
		// panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		// fmt.Fprintf(w, "%s\n", err.Error())
		return err
		// panic(err)
	}

	w.Write(content)
	return err
}

func main() {
	// 1. 错误的类型
	// const filename = "1.txt"
	// file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// 2. 错误的类型断言处理
	// 3. 自定义错误
	// err = errors.New("A customed error")
	// if err != nil {
	// 	// fmt.Printf("Error: {%T}, {%v}", err, err)
	// 	if pErr, ok := err.(*os.PathError); !ok {
	// 		panic(err)
	// 	} else {
	// 		fmt.Printf("%s,%s,%s,%s\n", pErr.Op, pErr.Path, pErr.Err, pErr.Error())
	// 	}
	// }

	// defer file.Close()

	// 练习：实现一个静态http服务器，并且在出错的时候优雅的处理
	// 错误处理的原则：第一，不要把内部的错误直接给到外部，一来不安全，二来会让普通用户疑惑
	//                第二，有统一的处理过程

	// 路由
	http.HandleFunc(txtPath, errWrapper(GetFile))

	// 启动服务器
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

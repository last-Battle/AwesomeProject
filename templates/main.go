package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//tpl, err := template.ParseFiles("./templates/hello.tpl")
	//if err != nil {
	//	fmt.Println("create templates error", err)
	//	return
	//}
	//
	//// 利用给定数据渲染模板，并将结果写入w
	//user := UserInfo{
	//	Name:   "小王子",
	//	Gender: "男",
	//	Age:    18,
	//}
	////tpl.Execute(w, "martin")
	//tpl.Execute(w, user)
	htmlByte, err := ioutil.ReadFile("./hello.tmpl")
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	// 使用user渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}

func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println(err)
	}
}

// TODO
// this sentence to solve the conflicts between vue and go template
//template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server fail, err:", err)
		return
	}
	//f1 := Friend{Fname: "xiaofang"}
	//f2 := Friend{Fname: "wugui"}
	//t := template.New("test")
	//t = template.Must(t.Parse(
	//	`hello {{.UserName}}!
	//		{{ range .Emails }}
	//		an email {{ . }}
	//		{{- end }}
	//		{{ with .Friends }}
	//		{{- range . }}
	//		my friend name is {{.Fname}}
	//		{{- end }}
	//		{{ end }}`))
	//p := Person{UserName: "longshuai",
	//	Emails:  []string{"a1@qq.com", "a2@gmail.com"},
	//	Friends: []*Friend{&f1, &f2}}
	//t.Execute(os.Stdout, p)

	//模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
	//模板文件中使用{{和}}包裹和标识需要传入的数据。
	//传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
	//除{{和}}包裹的内容外，其他内容均不做修改原样输出。

}

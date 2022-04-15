package main

import (
	common "awesomeProject/Code/project/commom"
	"awesomeProject/Code/project/services"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var sum int64 = 0
var productNum int64 = 0
var mutex sync.Mutex

// 1. 系统技术的需求分析
//    前端承担高负载
//    在高并发情况下秒杀系统的超卖问题
//    后端接口需要满足横向扩展

// 实现demo
// 2. 静态的HTML页面，MVC，不用：简单粗暴 HTML+CSS+jQ

// 3. go 语言搭建HTTP服务器

// 4. 配置MySQL的开发环境
// go get github.com/go-sql-driver/mysql

// 5. 从MySQL中读取数据

// 6. 改造：
//    6.1 秒杀活动的特点是，持续时间短，流量大，考虑部署单独的系统，甚至独立的域名，不要部署在同一个机房。
//    6.2 尽量重新设计商品详情页，页面内容尽量静态化，减少甚至避免服务器请求MySQL
//    6.3 前端页面要部署在CDN上，需要氪金
//    6.4 秒杀系统的限流，需要考虑高并发
//    6.5 一般来说，单台系统的30万QPS属于极限，如果预计流量更大，需要分布到多台机器上
//    6.6 如果成功下单，会写数据库需要用MQ缓存，直接写MySQL是扛不住的

func Get1Product() bool {
	mutex.Lock()
	defer mutex.Unlock()
	if sum < productNum {
		sum += 1
		fmt.Println(sum)
		return true
	}
	return false
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if Get1Product() {
		fmt.Fprintf(w, "{\"OK\":true}")
		return
	}
	fmt.Fprintf(w, "{\"OK\":false}")
	return
}

func main() {
	// 路由：支持静态页面访问
	http.Handle("/frontend/web/", http.StripPrefix("/frontend/web/", http.FileServer(http.Dir("frontend/web"))))

	// 路由：
	http.HandleFunc("/getProduct", GetProduct)

	// 初始化数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	productMgr := services.NewProductManager("product", db)
	products, err := productMgr.SelectAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products[0])
	// productNum = products[0].Count
	productNum = 5

	err = http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

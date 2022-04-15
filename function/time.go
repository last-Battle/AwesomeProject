package function

import (
	"fmt"
	"time"
)

func TIME() {
	//now := time.Now() //获取当前时间
	//fmt.Printf("current time:%v\n", now)
	//
	//year := now.Year()     //年
	//month := now.Month()   //月
	//day := now.Day()       //日
	//hour := now.Hour()     //小时
	//minute := now.Minute() //分钟
	//second := now.Second() //秒
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//time1stamp1 := now.Unix()
	//time2stamp2 := now.UnixNano()
	//timeObj := time.Unix(time1stamp1, 0)
	//year := timeObj.Year()

	//const (
	//	Nanoseacond Duration = 1
	//	Microsecond          = 1000 * Nanoseacond
	//)

	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}

	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
}

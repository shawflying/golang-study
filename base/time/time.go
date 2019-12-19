package main

import (
	"fmt"
	"time"
	//"math/rand"
	//"strconv"
	"strconv"
	"math/rand"
)

func main() {
	fmt.Println("时间控件的使用")
	fmt.Println("Demo all!")
	tody := time.Now().Weekday()
	fmt.Println(tody)
	fmt.Println("tody", time.Now())
	fmt.Println(time.Hour)
	fmt.Println(time.UnixDate)

	//时间戳
	cur := time.Now()
	timestamp := cur.UnixNano() / (1000000 * 1000)
	fmt.Println(timestamp) //获取时间戳

	fmt.Println("时间戳：---", time.Now().Unix())

	timeUnix1 := time.Now().Unix() //已知的时间戳
	formatTimeStr1 := time.Unix(timeUnix1, 0).Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr1) //打印结果：2017-04-11 13:30:39

	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("20060102150405")
	fmt.Println("wx_disney：" + formatTimeStr) //打印结果：2017-04-11 13:30:39

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println("随机数：" + strconv.Itoa(r.Intn(1000000000000)))
	}

	strtime := time.Now().Format("2006-01-02")
	fmt.Println("格式化时间：", strtime)
	if strtime < "2017-06-22 16:00:00" {
		fmt.Println("还在时间范围内")
	} else {
		fmt.Println("超过了,活动结束了")
	}

	fmt.Println("AAAAAAAAAAAAAA", time.Now().Format("2006-01-02 15:04:05"))

}

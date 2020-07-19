package main

import (
	"fmt"
	"time"
)

// 如果结构体定义为 time.Time，在传入string的时间时，可用下面的函数转换为go时间格式

// 字符串转时间
func ConverTime(ts string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	layTime := "2006-01-02T15:04:05+08:00"
	zeroTime, err := time.ParseInLocation(layTime, ts, loc)
	if err != nil {
		fmt.Println("str convertime fail")
		return time.Time{}
	}
	return zeroTime
}

func main() {
	t := "2020-07-07T17:07:07+08:00"
	fmt.Println(ConverTime(t))
}

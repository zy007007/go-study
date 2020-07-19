package main

import (
	"fmt"
	"time"
	"strings"
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


func ConverShortTime(ts string) string {
	tsArray := strings.Split(ts, " ")
	newts := strings.Join([]string{tsArray[0], "T", tsArray[1], "+08:00"}, "")
	return newts
}


func main() {
	t := "2020-07-07 17:07:07"
	fmt.Println(ConverTime(ConverShortTime(t)))
}

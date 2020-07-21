package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now() //获取当前时间
	t := now.Add(time.Minute * -1) // 获取上一分钟时间
	fmt.Println(t)
}

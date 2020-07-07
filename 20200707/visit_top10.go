// 从程序启动起，定时读取nginx access.log日志，统计ip，若出现则redis值增加1
// awk '{print $1}' access.log | sort | uniq -c | sort -nr
// ip 所属地
// curl -s --user-agent foobar https://ip.cn/index.php?ip=113.200.212.74 | grep '所在地理位置'| awk --F '>' '{print $9}' | awk -F '<' '{print $1}'

package main

import (
	"cmd"
	"fmt"
	"github.com/go-redis/redis/v8" 
)

type IpCount struct {
	Ip    string
	Times int
	Area  string
}

func ReadFile() {}



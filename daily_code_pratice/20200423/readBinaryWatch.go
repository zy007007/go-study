package main
//二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
//
//每个 LED 代表一个 0 或 1，最低位在右侧。
//例如，上面的二进制手表读取 “3:25”。
//
//给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
//
//案例:
//
//输入: n = 1
//返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
// 
//
//注意事项:
//
//输出的顺序没有要求。
//小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
//分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。

// 个人思路
// 题是看懂了，手表确实很靓
// 有点排列组合的味道

import "fmt"

// by comment 王小宁。双循环，i代表小时，j代表分钟，数一下i和j二进制形式中1的个数，如果和num一样，则为当前num可以表示的时间。
func readBinaryWatch(num int) []string {
	result := []string{}

	for i:=0; i<12; i++ {
		for j:=0; j<60; j++ {
			b1 := fmt.Sprintf("%b", i)
			b2 := fmt.Sprintf("%b", j)
			sumOne := strings.Count(b1, "1") + strings.Count(b2, "1")
			if sumOne == num {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}

func main() {
	fmt.Println("vim-go")
}
// 收获：哦直接是用 %b 转化成了二进制，就是暴力解法

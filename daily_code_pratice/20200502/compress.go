package main
//给定一组字符，使用原地算法将其压缩。
//
//压缩后的长度必须始终小于或等于原数组长度。
//
//数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。
//
//在完成原地修改输入数组后，返回数组的新长度。
//
//进阶：
//你能否仅使用O(1) 空间解决问题？

// 个人思路
// 万物皆可map
// 知识点
// 1.map存值
// 2.字符数字转换


import "fmt"

// by comment golang
func compress(chars []byte) int {
	cLen := len(chars)
	if cLen == 0 || cLen == 1 {
		return cLen
	}

	read, fp, fpTime := 0, 0, 0
	b := chars[0:0]

	for read <= cLen-1 {
		if read + 1 == cLen || chars[read] != chars[read+1] {
			b = append(b, chars[fp:fp+1]...)

			if fpTime > 1 {
				b = append(b, []byte(strconv.Itoa(fpTime))...)
			}

			fp = read + 1
			fpTime = 0
		} else {
			if fpTime == 0 {
				fp = read
				fpTime = 2
			} else {
				fpTime ++ 
			}
		} 
		read ++
	}
	return len(b)
}

func main() {
	fmt.Println("vim-go")
}
// 收获：数字转字符数组，其他逻辑和原因还看不懂着

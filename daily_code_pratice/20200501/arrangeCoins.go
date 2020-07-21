package main
//你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
//
//给定一个数字 n，找出可形成完整阶梯行的总行数。
//
//n 是一个非负整数，并且在32位有符号整型的范围内。
//
// 个人思路
// 二分法，找到中间
// 二分个什么
// 靠数学走天下


import "fmt"


// by comment L.Ljubav.San
func arrangeCoins(n int) int {
	i := 1 
	for ;i<=n; i++ {
		n -= i
	} 
	return i-1
}

// by comment Gwalker
func arrangeCoins(n int) int {
	return int((-1 + math.Sqrt(float64(1+n*8)))/2)
}




func main() {
	fmt.Println("vim-go")
}

package main

//给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
//
//示例 1:
//
//输入: 16
//输出: true
//示例 2:
//
//输入: 5
//输出: false
//进阶：
//你能不使用循环或者递归来完成本题吗？

// 个人思路
// 直接用了4进制判断法

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}

	ns := strconv.FormatInt(int64(num), 4)
	return ns[0:1] == "1" && strings.Count(ns, "0") == len(ns)-1
}

// 收获：3，4幂次都这样子判断了，后面是不是都一样可以

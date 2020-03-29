package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

//个人思路
//１选最短的字符串长度
//２从０－length 开始判断n个字符 str[n] 是否相等
//３不用选最短了，还需要判断，直接２，对即将越界的字符进行判断,所以关键是：go语言如何判断数组越界，好像还是要选出最短

//以上思路不对，还是注释掉，决定使用类似类似冒泡排序
//测试用例均通过，提交报错

//res = res + string(strs[i][k])
//panic: runtime error: index out of range [0] with length 0
//main.longestCommonPrefix(0xc00004a100, 0x1, 0x4, 0x4c8a00, 0xc00000c200)
//solution.go, line 14
//main.__helper__(...)
//solution.go, line 29
//main.main()
//solution.go, line 57

import (
	"fmt"
	//	"sort"
)

func longestCommonPrefix(strs []string) string {
	//	str_to_length := make(map[string]int)
	//
	//	for i := 0; i < len(strs); i++ {
	//		str_to_length[strs[i]] = len(strs[i]) // {"flower":6, "flow":4, ...}
	//	}
	//
	//	var lens []int
	//	for _, k := range str_to_length {
	//		lens = append(lens, k)
	//	}
	//
	//	sort.Ints(lens)
	//
	//	min_len := lens[0] // 获得最小的长度，将所有字符串先减少到min_len长度
	//
	//	var new_strs []string
	//
	//	for k := 0; k < len(strs); k++ {
	//		new_data := string(strs[k][0:min_len]
	//	}
	var res string
	i := 0     // 数组第一个元素开始
	k := 0     // 字符串第一个字符
	count := 1 //相同字符串前缀的次数

	for i < len(strs) {
		for j := i + 1; j < len(strs); j++ {
			if string(strs[i][k]) == string(strs[j][k]) {
				count += 1
			}
		}
		if count == len(strs) {
			res = res + string(strs[i][k])
			k += 1
			count = 1
		} else {
			break
		}
	}

	return res

}

func main() {
	var nums int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &nums)

	var arrs []string
	fmt.Println("依次输入字符串，按空格分开")
	for i := 0; i < nums; i++ {
		var s string
		fmt.Scanf("%s", &s)
		arrs = append(arrs, s)
	}

	fmt.Printf("数组")
	fmt.Println(arrs)
	fmt.Printf("输出:")
	fmt.Println(longestCommonPrefix(arrs))
}

// 收获：一开始想找其他方法，最后还是先用普通的循环判断解决，但是系提交统代码错误，没搞懂中
//	进一步熟悉了字符串相关判断

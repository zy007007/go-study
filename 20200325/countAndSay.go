package main

//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
//注意：整数序列中的每一项将表示为一个字符串。
//
//
//
//示例 1:
//
//输入: 1
//输出: "1"
//解释：这是一个基本样例。
//示例 2:
//
//输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。

//个人思路
//找规律吗
// n=5 ，1  读作 111    3个1 2个2 一个1
// 结果还是看了评论才看懂题的意思--!
// 6=312211
// 但还是需要根据n来判断结果吧
// 如果和n没有关系，那么就只能将1-30，从1开始对应的值，存放到map中了
//   考一个重复子串长度 或 数组重复元素长度
//   数字字符串的疯狂转换
//   初始数组 a = [1]
//   tmp 数组 存放连续重复元素x
//        新数组存放  b = len(tmp) + x
//  map 保存数据

//自己写的，测试用例通过，耗时和内存有点高
//学习较快的代码，见 countAndSayV1

import (
	"fmt"
	"strconv"
)

func countAndSayV1(n int) string {
	dp := make([]string, n)
	sb := strings.Builder{} // 学习 strings  多类型写入方法
	dp[0] = "1"
	for i := 1; i < n; i++ {
		count := 1
		last := dp[i-1]
		for j := 0; j < len(last)-1; j++ {
			if last[j] == last[j+1] {
				count++
				continue // 学习continue的使用
			}
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(last[j])
			count = 1
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(last[len(last)-1])
		dp[i] = sb.String()
		sb.Reset()
	}
	return dp[n-1]
}

func countAndSay(n int) string {

	data := make(map[int]string)
	data[1] = "1"
	data[2] = "11"
	data[3] = "21"
	data[4] = "1211"

	seq := 5

	arr := []int{1, 2, 1, 1}

	for seq <= 30 {
		i := 0
		var tmp []int
		counts := 1
		j := i + 1
		for j < len(arr) {

			if arr[i] == arr[j] {
				counts += 1
				j += 1
			} else {
				tmp = append(tmp, counts)
				tmp = append(tmp, arr[i])
				i = j
				j += 1
				counts = 1
			}
		}
		tmp = append(tmp, counts)
		tmp = append(tmp, arr[i])

		var str string
		for _, v := range tmp {
			str += strconv.Itoa(v)
		}

		data[seq] = str

		seq += 1

		arr = tmp

	}

	return data[n]
}

func main() {
	var num int
	fmt.Printf("输入:")
	fmt.Scanf("%d", &num)

	fmt.Printf("输出:")
	fmt.Println(countAndSay(num))
}

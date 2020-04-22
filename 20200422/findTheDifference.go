package main

//给定两个字符串 s 和 t，它们只包含小写字母。
//
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//
//
//
//示例:
//
//输入：
//s = "abcd"
//t = "abcde"
//
//输出：
//e
//
//解释：
//'e' 是那个被添加的字母。

// 个人思路
// map 可以解决

func findTheDifference(s string, t string) byte {
	dict := make(map[byte]int)

	for _, k := range []byte(t) {
		dict[k] = 1
	}

	for _, k := range []byte(s) {
		if _, ok := dict[k]; !ok {
			return k
		}
	}
	return nil
}

package main
//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
//示例 1:
//
//输入: "hello"
//输出: "holle"
//示例 2:
//
//输入: "leetcode"
//输出: "leotcede"
//说明:
//元音字母不包含字母"y"。

// 个人思路
// 元音字母：aeiou
// 前后一起遍历，判断，若是则交换


func reverseVowels(s string) string {
	res := []byte(s)
	match := map[byte]int{
		'a':1, 'e':1, 'i':1, 'o':1, 'u':1,
		'A':1, 'E':1, 'I':1, 'O':1, 'U':1,
	}
	left, right := 0, len(res)-1

	for left < right {
		l, r := match[res[left]], match[res[right]]
		if l == 1 && r == 1 {
			res[left], res[right] = res[right], res[left]
			left ++
			right --
		} else if l == 0 {
			left ++
		} else if r == 0 {
			right --
		} else {
			left ++
			right --
		}
	}

	return string(res)

}
// 收获： 字符串转字符数组 []byte()   字符数组转字符串string([]byte())

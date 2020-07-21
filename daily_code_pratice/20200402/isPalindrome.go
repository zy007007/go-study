package main

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false

// 个人思路
// 关键在如何忽略其他字符
// 笨办法,字母数字放入新的切片，判断即可
// 有unicode.IsDigit   unicode.IsLetter

func isPalindrome(s string) bool {
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	i, j := 0, len(s)-1

	for i < j {
		left, right := rune(s[i]), rune(s[j])

		if !isValid(left) && !isValid(right) {
			i++
			j--
		} else if !isValid(left) {
			i++
		} else if !isValid(right) {
			j--
		} else if unicode.ToUpper(left) != unicode.ToUpper(right) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

// 收获：https://golang.org/pkg/unicode/#ToUpper unicode 这个包，判断rune类型是否为数字、字母, ToUpper变大写

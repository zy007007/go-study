package main

import "fmt"

func main() {
	var strss = []string{
		"qwr",
		"234",
		"yui",
	}
	var strss2 = []string{
		"qqq",
		"aaa",
		"zzz",
		"zzz",
	}
	strss = append(strss, strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss)
	fmt.Println("vim-go")
}

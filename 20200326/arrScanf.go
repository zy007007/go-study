package main

import "fmt"

func main() {
	var length int
	var arr []int
	fmt.Printf("数组长度:")
	fmt.Scanf("%d", &length)

	fmt.Printf("输入元素:")
	for i := 0; i < length; i++ {
		var data int
		fmt.Scanf("%d", &data)

		arr = append(arr, data)
	}

}

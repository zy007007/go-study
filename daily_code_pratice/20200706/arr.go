package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)
}

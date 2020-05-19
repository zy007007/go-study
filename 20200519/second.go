package main

import "fmt"

func secondMax(arr []int) int {
	var res int
	max1, max2 := arr[0], arr[1]
	for i := 2; i < len(arr)-2; i++ {
		if max1 < arr[i] {
			max1 = arr[i]
		}

		if max2 < arr[i+1] {
			max2 = arr[i+1]
		}

		if max1 > max2 {
			res = max2
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}

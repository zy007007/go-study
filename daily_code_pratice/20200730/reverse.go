package main

import (
    "fmt"
)

func main() {
	var arr  [7]int = [7]int{1,2,4,6,12,2234,432111}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(arr *[7]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

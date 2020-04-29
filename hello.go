package main

import (
	"fmt"
	"math"
)

func main() {
	test := "zhaoyi3"
	test1 := []int{1, 2, 3, 4}
	fmt.Println("vim-go")
	fmt.Println(string(test[0:1]))

	fmt.Printf("test1[0] address:")
	fmt.Println(&(test1[0]))

	//fmt.Printf("test1[0:1] address:")
	//fmt.Println(&(test1[0:1]))

	fmt.Printf("test1[0:1][0] address:")
	fmt.Println(&(test1[0:1][0]))

	fmt.Println(1 / 2)

	fmt.Println(math.MinInt32 - 1)

}

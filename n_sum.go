package main

import (
	"fmt"
)

func main() {
	var num int
	sum := 0
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}

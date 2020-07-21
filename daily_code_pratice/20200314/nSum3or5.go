package main

import (
	"fmt"
)

func IsNumOk(thisnum, flag int) bool {
	if thisnum%flag == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var num int
	sum := 0

	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		if IsNumOk(i, 3) || IsNumOk(i, 5) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println("sum=", sum)
}

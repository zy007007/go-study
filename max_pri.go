package main

import (
	"fmt"
	"math"
)

func IsSu(n int) bool {
	flag := true
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			flag = false
		}
	}
	return flag
}

func main() {
	var num int
	num = 9223372036854775807
	for i := 3; i <= num; i++ {
		if IsSu(i) {
			fmt.Printf("%d ", i)
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandNum() bool {
	rand.Seed(time.Now().UnixNano())
	tmp := 1 * rand.Float64()

	res, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", tmp), 64)
	fmt.Println("rand is ", res)

	if res > 0.5 {
		return true
	} else {
		return false
	}
}

func SumNum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func MultNum(n int) int {
	mult := 1
	for i := 1; i <= n; i++ {
		mult *= i
	}
	return mult
}

func main() {
	var res int
	var des string
	var num int
	fmt.Scanf("%d", &num)

	if RandNum() {
		res, des = SumNum(num), "sum"
	} else {
		res, des = MultNum(num), "mult"
	}

	fmt.Println("result is ", des, " ", res)
}

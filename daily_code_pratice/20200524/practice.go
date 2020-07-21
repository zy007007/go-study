package main

import (
	"fmt"
)

func main() {
	for i, n := 100, 0; i < 201; i++ {
		sqt := int(math.Sqrt(float64(i + 1)))
		for j := 2; j <= sqt; j++ {
			if i%j == 0 {
				break
			}
		}
		n++
		if n%10 == 0 {
			fmt.Println()
		}
	}

}

func fibo(n int) int {
	if n < 3 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)

	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	n1, n2 := 1, 1
	for i := 3; i < n; i++ {
		n1, n2 = n1+n2, n1
	}
	return
}

func cn(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	sqt := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqt; i++ {
		if num%i == 0 {
			sum += i + (num / i)
		}
	}
	return num == sum
}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		split[i] = reverse(split[i])
	}
	return strings.Join(split, " ")
}

func reverse(s string) string {
	tmp := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	return string(tmp)
}

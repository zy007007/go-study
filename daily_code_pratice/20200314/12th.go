package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 12; i++ {
		for j := 1; j <= i; j++ {
			res := i * j
			desc := strconv.Itoa(i) + "*" + strconv.Itoa(j) + "="
			fmt.Printf("%s%d%s", desc, res, "\t")
		}
		fmt.Println()
	}
}

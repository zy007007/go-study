package main

import (
	"fmt"
)


func main() {
	data := make(map[string][]int)

	ids := []int{1,2,3,4,5,6}
	for _, v := range ids {
		if v % 2 == 0 {
			data["oushu"] = append(data["oushu"], v)
		}
	}
	fmt.Println(data)
}

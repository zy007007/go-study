package arrTools

import (
	"fmt"
)

// 输入n长度的整形数组
func arrInput(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		var data int
		fmt.Scanf("%d", &data)
		arr = append(arr, data)
	}
	return arr
}

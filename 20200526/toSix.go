package main

import (
	"fmt"
)

// 10240 转换未十六进制
func toSix(n int) string {
	result := ""
	h := map[int]string{0:"0", 1:"1", 2:"2"...,10:"a", 11:"b"...,15:"f"}
	for ; n>0; n /= 16 {
		tmp := h[n%16]
		result = result + tmp
	}
	return result
}


func main() {
	fmt.Println("vim-go")
}

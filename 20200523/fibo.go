package main

import "fmt"



func fi1(n int) int{
	if n == 0 {
		return 0
	} else if  n == 1 {
		return 1
	} else if n > 1 {
		return fi1(n-1) + fi1(n-2)
	} else {
		return -1
	}
}



func main() {
	fmt.Println("vim-go")
}

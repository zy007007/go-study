package main

import (
	"fmt"
)

func main() {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break // break 但是go func 的ch还处在发送的状态，无法结束
		}
	}
}

package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}

// 预期输出200
// 淦，简单粗暴的预期错了
// 1+...+100=5050
// 5050*2 = 10100

// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(2) // 把示例代码的１修改为２，可以看见输出的内容发生了极大的变化
			      // 可以进一步了解到，当设置为１个逻辑处理器时，调度器会自己进行goroutine之间的切换
			      //　当设置２个逻辑处理器时，每个goroutine　运行在单个逻辑处理器上面

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

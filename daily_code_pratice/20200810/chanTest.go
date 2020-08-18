package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int, isTrue chan<- bool) {
	for j := range jobs {
		fmt.Printf("I am %d working at jobs %d\n", id, j)
		results <- j * 2
		isTrue <- true
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	isTrue := make(chan bool)
	for i := 1; i < 4; i++ {
		go worker(i, jobs, results, isTrue)
	}
	for k := 1; k < 10; k++ {
		jobs <- k
		<-isTrue
	}
	for v := 1; v < 10; v++ {
		<-results
	}
}

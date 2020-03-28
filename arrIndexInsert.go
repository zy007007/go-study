package main

import "fmt"

func insertArrindex(a []int, index int, data int) {
	//var tmp []int
	//tmp = a[index:]
	tmp := append([]int{}, a[index:]...)
	fmt.Println("tmp", res)
	a = append(a[0:index], data)
	fmt.Println(a)
	fmt.Println(res)
	a = append(a, res...)
}

func InsertStringSliceCopy(slice, insertion []int, index int) []int {
	result := make([]int, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func main() {
	arr := []int{1, 2, 3}
	//arr1 := []int{1, 2, 3}
	//arr2 := []int{1, 2, 3}
	//fmt.Println(InsertStringSliceCopy(arr, []int{0}, 0))
	//fmt.Println(InsertStringSliceCopy(arr1, []int{-2}, 1))
	//fmt.Println(InsertStringSliceCopy(arr2, []int{-3}, 3))

	insertArrindex(arr, 0, 0)
	fmt.Println(arr)
}

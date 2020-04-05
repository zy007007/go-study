package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr=", arr)

	fmt.Println("arr[2:4]=", arr[2:4])

	fmt.Println("arr[len(arr)-1:] 尾元素:", arr[len(arr)-1:])

	fmt.Println("arr[len(arr)-2:] 两尾元素:", arr[len(arr)-2:])

	fmt.Println("arr[:len(arr)-1] 不包含尾元素，或者可以称为删除尾元素", arr[:len(arr)-1])

	fmt.Println("arr[1:] 不包含头元素，或者可以称为删除头元素", arr[1:])

	arr = append(arr[0:3], arr[4:len(arr)]...)
	fmt.Println("append(arr[0:3], arr[4:len(arr)]...) 删除下标为k=3元素即4", arr)

	arr0 := append([]int{0}, arr...)
	fmt.Println("append([]int{0}, arr...) 头插一个元素", arr0)

	arr1 := append(arr0, arr...)
	fmt.Println("append(arr0, arr...) 头插一个切片", arr1)

	arr2 := append(arr, arr0...)
	fmt.Println("append(arr, arr0...) 尾插一个切片", arr2)

	//Output
	//arr= [1 2 3 4 5 6 7]
	//arr[2:4]= [3 4]
	//arr[len(arr)-1:] 尾元素: [7]
	//arr[len(arr)-2:] 两尾元素: [6 7]
	//arr[:len(arr)-1] 不包含尾元素，或者可以称为删除尾元素 [1 2 3 4 5 6]
	//arr[1:] 不包含头元素，或者可以称为删除头元素 [2 3 4 5 6 7]
	//append(arr[0:3], arr[4:len(arr)]...) 删除下标为k=3元素即4 [1 2 3 5 6 7]
	//append([]int{0}, arr...) 头插一个元素 [0 1 2 3 5 6 7]
	//append(arr0, arr...) 头插一个切片 [0 1 2 3 5 6 7 1 2 3 5 6 7]
	//append(arr, arr0...) 尾插一个切片 [1 2 3 5 6 7 0 1 2 3 5 6 7]

	//后续有新的get点，及时补充。关键字：切片操作
}

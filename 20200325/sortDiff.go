package main

// 代码来自 https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/
// 用于测试效果

import (
	"fmt"
	"sort"
)

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0, 10}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

}

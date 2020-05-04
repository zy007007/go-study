package main

//给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
//
//找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
//
//示例:
//
//输入:
//[[0,0],[1,0],[2,0]]
//
//输出:
//2
//
//解释:
//两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

// 个人思路
// 读懂了，两点之间的距离

import "fmt"

// by comment ElliotXX
func dist(i, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	hash := map[int]int{}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				d := dist(point[i], point[j])

				if hash[d] > 0 {
					ans += hash[d] * 2
				}
				hash[d]++
			}
		}
		hash = map[int]int{}
	}
	return ans

}

func main() {
	fmt.Println("vim-go")
}

// 收获：*2属实没看懂；懂了，因为如果flag有第三个相同的了，那么也和第二个相同

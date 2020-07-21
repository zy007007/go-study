package main

//给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
//
//网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
//
//岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

// 个人思路
// 确定什么情况下 1 = 3 cm / 1 = 2 cm / 1 = 1 cm / 1 = 4 cm
// 1 上下左右4个0，l=4
// 1 	     3个0，l=3
// 1         2个0，l=2
// 1         1个0，l=1

import "fmt"

// by comment 天驰🎩
func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					res -= 2
				}

				if j-1 >= 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}

// 收获: -2 这个就很妙啊

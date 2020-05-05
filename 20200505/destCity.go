package main

//给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi 直接前往 cityBi 。请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。
//
//题目数据保证线路图会形成一条不存在循环的线路，因此只会有一个旅行终点站。
//
//
//示例 1：
//
//输入：paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
//输出："Sao Paulo"
//解释：从 "London" 出发，最后抵达终点站 "Sao Paulo" 。本次旅行的路线是 "London" -> "New York" -> "Lima" -> "Sao Paulo" 。

import "fmt"

// by comment 梦想追风
func destCity(paths [][]string) string {
	length := len(paths)
	if length == 1 {
		return paths[0][1]
	}

	pathsMap := make(map[string]string, 10)
	ok := false

	for i := 0; i < length; i++ {
		pathsMap[paths[i][0]] = paths[i][1]
	}

	index := paths[0][0]
	for {
		if _, ok := pathsMap[index]; ok {
			index = pathsMap[index]
		} else {
			return index
		}
	}
	return paths[0][1]
}

func main() {
	fmt.Println("vim-go")
}

// 收获：使用了map，最终就是：第二位在第一位中不存在

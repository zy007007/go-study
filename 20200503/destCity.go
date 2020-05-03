package main

//给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi 直接前往 cityBi 。请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。
//
//题目数据保证线路图会形成一条不存在循环的线路，因此只会有一个旅行终点站。

// 个人思路
// 居然想不到怎么解。。。
// 把起点放入一个数组，把终点放入一个数组
// 遍历终点的每一个元素，如果不在起点，则返回

import "fmt"

func destCity(paths [][]string) string {

	var start []string
	var end []string

	for _, v := range path {
		start = append(start, v[0])
		end = append(end, v[1])
	}

	var res string
	for _, e := range end {
		if !isExist(e, start) {
			res = e
		}
	}
	return res
}

func isExist(s string, arr []string) bool {
	var res bool
	for _, v := range arr {
		if v == res {
			res = true
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}

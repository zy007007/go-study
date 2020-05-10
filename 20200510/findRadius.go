package main
//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
//现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
//
//所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。
//
//说明:
//
//给出的房屋和供暖器的数目是非负数且不会超过 25000。
//给出的房屋和供暖器的位置均是非负数且不会超过10^9。
//只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
//所有供暖器都遵循你的半径标准，加热的半径也一样。
//示例 1:
//
//输入: [1,2,3],[2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。


// by comment 真小雨
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	low, high := 0, 1000000000
	for low < high {
		mid := low + (high-low)/2
		hIdx := 0
		left, right := heaters[hIdx]-mid, heaters[hIdx]+mid
		var i int
		for i=0; i<len(houses); {
			if houses[i] < left {
				break
			} else if houses[i] > right {
				hIdx ++
				if hIdx > = len(heaters){
					break
				}
				left, right = heaters[hIdx]-mid, heaters[hIdx]+mid
				continue
			}
			i++
		}
		if i == len(houses) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

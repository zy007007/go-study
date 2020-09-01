func judgeCircle(moves string) bool {
	var dic = make(map[string]int)
	dic["R"] = 1
	dic["L"] = -1
	dic["U"] = 2
	dic["D"] = -2

	var count int
	for _, v := range moves {
		count += dic[v]
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}

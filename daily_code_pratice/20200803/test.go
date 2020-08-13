func strs(s1, s2 string) int {
	if len(s2) == 0 {
		return 0
	}

	var i, j int
	for i := 0; i < len(s1)-len(s2) + 1 ; i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i+j] != s2[j] {
				break
			}
		}

		if len(s2) == j {
			return i
		}
	}
	return -1
}

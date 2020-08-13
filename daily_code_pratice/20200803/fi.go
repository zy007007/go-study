



func fi(n int) []int {
	res := make([]int, n, n)
	for i:=0; i<n; i++ {
		if i <= 1 {
			res[i] = 1
		} else {
			res [i] = res[i-1] + res[i-2]
		}
	}
	return res
}

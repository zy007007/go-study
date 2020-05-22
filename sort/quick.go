package main

func quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}

	return append(append(quick(left), v), quick(right)...)
}

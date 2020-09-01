func fairCandySwap(A []int, B []int) []int {
	sum = sum(A) + sum(B)

	half = sum / 2

	a_chazhi = sum(A) - half

	b_chazhi = sum(B) - half

	for _, a := range A {
		for _, b := range B {
			chazhi = a - b
			if a_chazhi == chazhi || b_chazhi == chazhi {
				return [a,b]
			}
		}
	}


}



func remove(arr []int) []int {
	for i := range arr {
		if i > len(arr) -1 {
			return arr
		}

		if i < len(arr) - 1 && arr[i] == arr[i+1] {
			cpoy(arr[i:], arr[i+1:])
			arr = arr[:len(arr)-1]
			return remove(arr)
		}
	}
}



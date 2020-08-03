


func Sort(arr []int) {
	for i:=0; i<len(arr); i++ {
		for j := i; j<len(arr); j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

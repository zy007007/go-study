package main

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		pos := i
		for j := i + 1; i < len(arr); j++ {
			if arr[pos] > arr[j] {
				pos = j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i]
	}
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

package main

func selectSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		pivot := i
		pivotKey := a[i]
		for j := i + 1; j < n; j++ {
			if a[j] < pivotKey {
				pivot = j
				pivotKey = arr[j]
			}
		}
		a[pivot], a[i] = a[i], a[pivot]
	}
}

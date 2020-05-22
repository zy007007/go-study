package main

func insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1

		if j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}

		if j+1 == i {
			continue
		}

		arr[j+1] = val
	}
}

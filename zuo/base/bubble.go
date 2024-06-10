package base

func bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
		}
	}

	return arr
}

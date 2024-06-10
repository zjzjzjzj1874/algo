package base

// 调用方：请保证传入的l <= r
func mergeSort(arr []int, l, r int) []int {
	if l == r {
		return arr
	}
	m := l + (r-l)>>1 // 取两数的中点

	// 归并排序的思想：分治，把数组使用二分，分成最小的数组来俩俩比较大小
	mergeSort(arr, l, m)   // 拆分前半部分
	mergeSort(arr, m+1, r) // 拆分后半部分
	merge(arr, l, m, r)    // 对数组的前后两部分使用归并排序

	return arr
}

func merge(arr []int, l, m, r int) []int {
	temp := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			temp[i] = arr[p1]
			p1++
		} else {
			temp[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		temp[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= r {
		temp[i] = arr[p2]
		i++
		p2++
	}

	for i := 0; i < len(temp); i++ {
		arr[l+i] = temp[i]
	}

	return arr
}

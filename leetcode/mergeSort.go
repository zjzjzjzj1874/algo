package leetcode

import "fmt"

// MergeSort 归并排序
func MergeSort(nums []int) {
	fmt.Println("排序前：", nums)
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println("排序后：", nums)
}

func mergeSort(nums []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	merge(nums, l, mid, r)
}

func merge(nums []int, l, m, r int) {
	helper := make([]int, r-l+1)
	p1 := l
	p2 := m + 1
	i := 0
	for p1 <= m && p2 <= r {
		if nums[p1] > nums[p2] {
			helper[i] = nums[p2]
			p2++
		} else {
			helper[i] = nums[p1]
			p1++
		}
		i++
	}

	for p1 <= m {
		helper[i] = nums[p1]
		p1++
		i++
	}
	for p2 <= r {
		helper[i] = nums[p2]
		p2++
		i++
	}

	for i := 0; i < len(helper); i++ {
		nums[l+i] = helper[i]
	}
}

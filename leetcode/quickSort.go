package leetcode

import "fmt"

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	quickHelper(nums, 0, len(nums)-1)

	fmt.Println(nums)
}

// 辅助函数，(l,r)表示nums[l:r]需要排序的；
func quickHelper(nums []int, l, r int) {
	if l < r {
		sliptIdx := partition(nums, l, r)
		quickHelper(nums, l, sliptIdx-1) // 不管等不等于，假设不等于，取前一个即可
		quickHelper(nums, sliptIdx+1, r)
	}
}

func partition(nums []int, l, r int) int {
	base := nums[r] // 以最后一个元素作为基准排序值

	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < base { // 这里，只有当j对应的值小于base值时，才对ij交换。i维护的是一条小于base的基准线
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[r] = nums[r], nums[i+1] // 这里是对基准线后面的值交换

	return i + 1
}

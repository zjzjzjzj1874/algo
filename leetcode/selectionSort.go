package leetcode

import "fmt"

// SelectionSort 选择排序
func SelectionSort(nums []int) {
	fmt.Println("排序前：", nums)
	for i := 0; i < len(nums); i++ {
		minIdx := i // 最小数的下标

		// 每次循环，搞定一个最小的数
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				nums[minIdx], nums[j] = nums[j], nums[minIdx]
			}
		}
	}

	fmt.Println("排序后：", nums)
}

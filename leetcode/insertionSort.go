package leetcode

import "fmt"

// InsertionSort 插入排序：从小到大，依次做到有序
func InsertionSort(nums []int) {
	fmt.Println("排序前：", nums)
	for i := 1; i < len(nums); i++ { // 做到从0~i之间有序
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println("排序后：", nums)
}

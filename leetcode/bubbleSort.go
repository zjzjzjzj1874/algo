package leetcode

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	fmt.Println("排序前：", nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 冒泡排序也是，把小的放前面，每一轮搞定一个最小的数
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	fmt.Println("排序后：", nums)
}

package leetcode

// Leetcode.27  => 原地移除元素

func removeElement(nums []int, val int) int {
	res := len(nums)

	for i := 0; i < res; {
		if val != nums[i] {
			i++
			continue
		}

		nums[i], nums[res-1] = nums[res-1], nums[i] // 交换位置，换到最后面，然后res--
		res--
	}

	return res
}

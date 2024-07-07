package leetcode

// GetMax 递归法求最大值
func GetMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return getMax(nums, 0, len(nums)-1)
}

func getMax(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}

	mid := l + (r-l)/2
	left := getMax(nums, l, mid)
	right := getMax(nums, mid+1, r)

	if left > right {
		return left
	} else {
		return right
	}
}

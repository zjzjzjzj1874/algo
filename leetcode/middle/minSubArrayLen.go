package middle

import "math"

// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的
// 子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 提示：
//
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105
//
// 进阶：
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

// 解题：滑动窗口，注意左右指针的边界条件以及判断的逻辑
func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt

	left := 0
	right := 0
	sum := 0

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if min > right-left+1 {
				min = right - left + 1
			}
			sum -= nums[left]
			left++ // 缩小长度
		}

		right++ // 扩大数组长度，即尝试增加sum
	}

	if min == math.MaxInt {
		min = 0
	}

	return min
}

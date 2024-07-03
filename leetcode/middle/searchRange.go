package middle

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
// 示例 2：
//
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
// 示例 3：
//
// 输入：nums = [], target = 0
// 输出：[-1,-1]
//
// 提示：
//
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums 是一个非递减数组
// -109 <= target <= 109

// 解题：二分法
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l := 0
	r := len(nums) - 1
	first := -1
	last := -1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			first = mid
			r = mid - 1 // 重点:第一次等于，往左边去找
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	l = 0
	r = len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			last = mid
			l = mid + 1 // 重点，第二次，往右边去找
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{first, last}
}

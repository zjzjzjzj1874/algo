package dichotomy

// 704. 二分查找
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
// 示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
// 输出: -1
// 解释: 2 不存在 nums 中因此返回 -1
//
// 提示：
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。

// 解题：二分法
func search(nums []int, target int) int {
	n := len(nums)

	left := -1
	right := n

	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if right == n || nums[right] != target {
		return -1
	}

	return right
}

// 左闭右开的写法
func searchV1(nums []int, target int) int {
	n := len(nums)
	l := 0 // 左闭
	r := n // 右开

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target { // target在左半区
			r = mid
		} else {
			l = mid + 1 // +1是因为闭区间，不重复取l，所以mid++
		}
	}

	return -1
}

// 解题：左闭右闭的写法
func searchV2(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// 解题：双开的写法
func searchV3(nums []int, target int) int {
	n := len(nums)
	l := -1
	r := n
	for l+1 < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}

	return -1
}

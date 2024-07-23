package dichotomy

// LCR 070. 有序数组中的单一元素
// 给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。
//
// 示例 1:
//
// 输入: nums = [1,1,2,3,3,4,4,8,8]
// 输出: 2
// 示例 2:
//
// 输入: nums =  [3,3,7,7,10,11,11]
// 输出: 10
//
// 提示:
//
// 1 <= nums.length <= 105
// 0 <= nums[i] <= 105
//
// 进阶: 采用的方案可以在 O(log n) 时间复杂度和 O(1) 空间复杂度中运行吗？
//
// 注意：本题与主站 540 题相同：https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		if mid%2 == 0 {
			if mid != 0 && nums[mid] == nums[mid-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if mid != n && nums[mid] == nums[mid+1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return nums[right]
}

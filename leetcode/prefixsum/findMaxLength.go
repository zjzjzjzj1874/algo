package prefixsum

// LCR 011. 连续数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
// 示例 1：
//
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
// 示例 2：
//
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
//
// 注意：本题与主站 525 题相同： https://leetcode-cn.com/problems/contiguous-array/
func findMaxLength(nums []int) (ans int) {
	// 前缀和
	sum := 0
	kMap := make(map[int]int) // map[sum]index => 即：记录当sum = 这个值时，最早出现的index是哪个。
	kMap[0] = -1

	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}

		if val, ok := kMap[sum]; ok {
			ans = max(ans, i-val)
		} else {
			kMap[sum] = i
		}
	}

	return
}

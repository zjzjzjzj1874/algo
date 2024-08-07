package prefixsum

// 930.和相同的二元子数组
// 560.和为K的子数组
// LCR010.和为K的子数组

// 给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
//
// 子数组 是数组的一段连续部分。
//
// 示例 1：
//
// 输入：nums = [1,0,1,0,1], goal = 2
// 输出：4
// 解释：
// 有 4 个满足题目要求的子数组：[1,0,1]、[1,0,1,0]、[0,1,0,1]、[1,0,1]
// 示例 2：
//
// 输入：nums = [0,0,0,0,0], goal = 0
// 输出：15
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// nums[i] 不是 0 就是 1
// 0 <= goal <= nums.length
func numSubArraysWithSum(nums []int, goal int) (ans int) {
	prefixSum := make(map[int]int) // map[sum]count
	prefixSum[0] = 0
	sum := 0
	for _, num := range nums {
		prefixSum[sum]++ // 这个放前面，就不用在前面特殊处理0：1这个键值对
		sum += num
		ans += prefixSum[sum-goal]
	}

	return
}

package backtrack

// LCR 102. 目标和
// 给定一个正整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
// 示例 1：
//
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 示例 2：
//
// 输入：nums = [1], target = 1
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
//
// 注意：本题与主站 494 题相同： https://leetcode-cn.com/problems/target-sum/

// 这是这道题的回溯解法，但是回溯效率较低，还有DP可供选择，后续使用DP解法来解
func findTargetSumWays(nums []int, target int) (ans int) {
	// 思考：回溯法是不是也可以解决这个问题
	n := len(nums)
	var dfs func(index, sum int)

	dfs = func(index, sum int) {
		if index == n {
			if sum == target {
				ans++
			}
			return
		}
		next := index + 1
		dfs(next, sum+nums[index])
		dfs(next, sum-nums[index])
	}

	// 从0开始，但是实际是到N，所以第一个才开始计数
	dfs(0, 0)

	return
}

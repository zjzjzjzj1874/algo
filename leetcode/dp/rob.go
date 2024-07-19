package dp

// LCR 090. 打家劫舍 II
// 一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
// 示例 1：
//
// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 示例 2：
//
// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
// 偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 3：
//
// 输入：nums = [0]
// 输出：0
//
// 提示：
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
// 注意：本题与主站 213 题相同： https://leetcode-cn.com/problems/house-robber-ii/
func rob(nums []int) int {
	// 动态规划
	// dp[0] = nums[0]
	// dp[1] = max(nums[0], nums[1])
	// dp[2] = max(dp[1], nums[2] + dp[0])
	// 状态转移方程
	// dp[i] = max(dp[i-1], nums[i] + dp[i-2])
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	if n == 3 {
		return max(nums[0], nums[1], nums[2])
	}

	// 这个问题有环，所以变成，nums[0:n-2],nums[1:n-1]之间的较大值
	var getMax func(start, end int) int
	getMax = func(start, end int) int {
		dp := make([]int, n-1)
		dp[0] = nums[start]
		dp[1] = max(nums[start], nums[start+1])

		for i := start + 2; i < end; i++ {
			// 这里-start是因为起点从start开始的
			dp[i-start] = max(dp[i-1-start], dp[i-2-start]+nums[i])
		}

		return dp[end-1-start]
	}

	return max(getMax(1, n), getMax(0, n-1))
}

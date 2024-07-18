package dp

// LCR 103. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
// 示例 4：
//
// 输入：coins = [1], amount = 1
// 输出：1
// 示例 5：
//
// 输入：coins = [1], amount = 2
// 输出：2
//
// 提示：
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// 注意：本题与主站 322 题相同： https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0 // base case，组成 0 金额的最少硬币数是 0

	for i := 0; i <= amount; i++ { // 遍历金额
		for j := 0; j < len(coins); j++ { // 遍历面额
			if i < coins[j] { // 当前金额小于硬币面值时，跳过本次循环
				continue
			}
			// 状态转移方程：这表示，组成金额i所需的最少硬币数可以通过取不使用 coins[j] 和使用 coins[j] 这两种情况中的较小值来更新。
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
			// 这句话的意义在于更新 dp[i] 的值，即组成金额 i 所需的最少硬币数。
			//dp[i - coins[j]] 表示当前金额 i 减去当前硬币面值 coins[j] 后的金额 i - coins[j] 所需的最少硬币数。
			//dp[i - coins[j]] + 1 表示在 dp[i - coins[j]] 的基础上再加上一个硬币 coins[j] 所需的硬币数。
			//min(dp[i], dp[i - coins[j]] + 1) 表示在不使用 coins[j] 和使用 coins[j] 这两种情况下，选择所需硬币数最少的一种情况来更新 dp[i] 的值。
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

package dp

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 示例 1：
//
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
// 示例 2：
//
// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶
//
// 提示：
//
// 1 <= n <= 45

// 解题：动态规划
// dp[0] = 0
// dp[1] = 1
// dp[2] = 2
// 状态转移方程
// dp[i] = dp[i-1]+dp[i-2]
// https://blog.csdn.net/fengzheng1232/article/details/127628014
// 对动态规划解法中状态转移方程f(n) = f(n-1) + f(n-2)的理解，这个说明稍微友好一点
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n)
	// 这里0不存在，我们统一把Index向前移动一步，这样不需要对dp数组进行扩容
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ { // 因为这里我们统一向前处理了一个
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// 解题：不使用数组，使用常量优化
func climbStairsWithOptimize(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	prePre := 1 // prePre = dp[0] = 1
	pre := 2    // pre = dp[1] = 2
	for i := 2; i < n; i++ {
		pre += prePre
		prePre = pre - prePre
	}

	return pre
}

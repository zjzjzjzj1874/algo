package dp

// LCR 099. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：一个机器人每次只能向下或者向右移动一步。
//
// 示例 1：
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 示例 2：
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// 注意：本题与主站 64 题相同： https://leetcode-cn.com/problems/minimum-path-sum/

// 解题：先用递归，再改记忆化表结构，最后使用dp
func minPathSum(grid [][]int) int {
	m := len(grid)    // grid的高度、row，行
	n := len(grid[0]) // grid的宽，col，列

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// LCR 166. 珠宝的最高价值
// 现有一个记作二维矩阵 frame 的珠宝架，其中 frame[i][j] 为该位置珠宝的价值。拿取珠宝的规则为：
//
// 只能从架子的左上角开始拿珠宝
// 每次可以移动到右侧或下侧的相邻位置
// 到达珠宝架子的右下角时，停止拿取
// 注意：珠宝的价值都是大于 0 的。除非这个架子上没有任何珠宝，比如 frame = [[0]]。
//
// 示例 1:
//
// 输入: frame = [[1,3,1],[1,5,1],[4,2,1]]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最高价值的珠宝
//
// 提示：
//
// 0 < frame.length <= 200
// 0 < frame[0].length <= 200

// 解题：这个题和上面一模一样，不过一个是求最大值，一个是求最小值
func jewelleryValue(frame [][]int) int {
	// 这个题和机器人走路一模一样的，不过一个求最大，一个求最小
	// 1  3  1
	// 1  5  1
	// 4  2  1

	// 转态转移：dp[m][n] = frame[m][n] + max(dp[m-1][n], dp[m][n-1])
	// dp[0] = 第一行的累加
	// dp[i][0] = 第一列的累加

	n := len(frame)    // 行 col
	m := len(frame[0]) // 列

	dp := make([][]int, n)

	dp[0] = make([]int, m)
	dp[0][0] = frame[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + frame[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = dp[i-1][0] + frame[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = frame[i][j] + max(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[n-1][m-1]
}

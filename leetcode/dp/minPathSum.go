package dp

import "math"

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
// 这个优化解题可以看一下：https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/solutions/2153802/jiao-ni-yi-bu-bu-si-kao-dpcong-hui-su-da-epvl/
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

// 931. 下降路径最小和
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
//
// 示例 1：
//
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 解释：如图所示，为和最小的两条下降路径
// 示例 2：
//
// 输入：matrix = [[-19,57],[-40,-5]]
// 输出：-59
// 解释：如图所示，为和最小的下降路径
//
// 提示：
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
func minFallingPathSum(matrix [][]int) int {
	// 走 i,j => i+1,j || i, j+1 || i+1, j+1
	// dp[0][i] = matrix[0][i]
	// dp[i][0] = dp[i-1][0] + matrix[i][0]
	// dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1], dp[i-1][j-1])

	n := len(matrix)    // 行，row
	m := len(matrix[0]) // 列， col
	dp := make([][]int, n)

	dp[0] = matrix[0]
	for i := 1; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1])
			} else if j < m-1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1], dp[i-1][j-1])
			} else {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	ans := dp[n-1][0]
	for i := 1; i < m; i++ {
		ans = min(ans, dp[n-1][i])
	}

	return ans
}

// 1289. 下降路径最小和 II
// 给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。
//
// 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
//
// 示例 1：
//
// 输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：13
// 解释：
// 所有非零偏移下降路径包括：
// [1,5,9], [1,5,7], [1,6,7], [1,6,8],
// [2,4,8], [2,4,9], [2,6,7], [2,6,8],
// [3,4,8], [3,4,9], [3,5,7], [3,5,9]
// 下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
// 示例 2：
//
// 输入：grid = [[7]]
// 输出：7
//
// 提示：
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// -99 <= grid[i][j] <= 99
func minFallingPathSumHard(grid [][]int) (ans int) {
	n := len(grid)
	// 还是只能dp，因为涉及到求min
	dp := make([][]int, n)
	dp[0] = grid[0]

	for i := 1; i < len(grid); i++ {
		dp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			// 求上一行不相邻元素最小值
			mini := math.MaxInt
			for k := 0; k < n; k++ {
				if j != k && dp[i-1][k] < mini {
					mini = dp[i-1][k]
				}
			}
			dp[i][j] = grid[i][j] + mini
		}
	}

	ans = math.MaxInt
	for i := 0; i < n; i++ {
		if ans > dp[n-1][i] {
			ans = dp[n-1][i]
		}
	}

	return ans
}

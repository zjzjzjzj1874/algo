package dp

// LCR 091. 粉刷房子
// 假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
// 当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。
//
// 例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
//
// 请计算出粉刷完所有房子最少的花费成本。
//
// 示例 1：
//
// 输入: costs = [[17,2,17],[16,16,5],[14,3,19]]
// 输出: 10
// 解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
// 最少花费: 2 + 5 + 3 = 10。
// 示例 2：
//
// 输入: costs = [[7,6,2]]
// 输出: 2
//
// 提示:
//
// costs.length == n
// costs[i].length == 3
// 1 <= n <= 100
// 1 <= costs[i][j] <= 20
//
// 注意：本题与主站 256 题相同：https://leetcode-cn.com/problems/paint-house/

// 解题：二维动态数组
func minCost(costs [][]int) int {
	n := len(costs) // 表示有n个房子
	if n == 0 {
		return 0
	}
	if n == 1 {
		return min(costs[0][0], min(costs[0][1], costs[0][2]))
	}
	// 本质上是三个条件中的两个互斥
	// 则使用二维数组来装
	// dp[0][j] = dp[0](红蓝绿)
	// dp[1][0] = dp[1]min(costs[1][1], costs[1][2])+dp[0][0]
	// dp[1][1] = dp[1]min(costs[1][0], costs[1][2])+dp[0][1]
	// dp[1][2] = dp[1]min(costs[1][0], costs[1][1])+dp[0][2]

	// dp[2][0] = dp[2]min(costs[2][1], costs[2][2])+dp[1][0]
	// dp[2][1] = dp[2]min(costs[2][0], costs[2][2])+dp[1][1]
	// dp[2][2] = dp[2]min(costs[2][0], costs[2][1])+dp[1][2]

	// 状态转移
	// dp[i][0] = dp[i]min(dp[i-1][1], dp[i-1][2])+costs[i][0]
	// dp[i][1] = dp[i]min(dp[i-1][0], dp[i-1][2])+costs[i][1]
	// dp[i][2] = dp[i]min(dp[i-1][0], dp[i-1][1])+costs[i][2]
	dp := make([][]int, n)
	dp[0] = []int{costs[0][0], costs[0][1], costs[0][2]}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = costs[i][j] + min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3])
		}
	}

	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func minCostWithOptimize(costs [][]int) int {
	n := len(costs) // 表示有n个房子
	if n == 0 {
		return 0
	}
	if n == 1 {
		return min(costs[0][0], min(costs[0][1], costs[0][2]))
	}
	// 本质上是三个条件中的两个互斥
	// 则使用二维数组来装
	// dp[0][j] = dp[0](红蓝绿)
	// dp[1][0] = dp[1]min(costs[1][1], costs[1][2])+dp[0][0]
	// dp[1][1] = dp[1]min(costs[1][0], costs[1][2])+dp[0][1]
	// dp[1][2] = dp[1]min(costs[1][0], costs[1][1])+dp[0][2]

	// dp[2][0] = dp[2]min(costs[2][1], costs[2][2])+dp[1][0]
	// dp[2][1] = dp[2]min(costs[2][0], costs[2][2])+dp[1][1]
	// dp[2][2] = dp[2]min(costs[2][0], costs[2][1])+dp[1][2]

	// 状态转移
	// dp[i][0] = dp[i]min(dp[i-1][1], dp[i-1][2])+costs[i][0]
	// dp[i][1] = dp[i]min(dp[i-1][0], dp[i-1][2])+costs[i][1]
	// dp[i][2] = dp[i]min(dp[i-1][0], dp[i-1][1])+costs[i][2]
	dp := make([][]int, n)
	dp[0] = []int{costs[0][0], costs[0][1], costs[0][2]}

	for i := 1; i < n; i++ {
		dp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = costs[i][j] + min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3])
		}
	}

	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

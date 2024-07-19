package dp

// 自拟题目：
// 给定一个数组coins和一个目标值target，数组值代表硬币的面额；
// 每个硬币只能选择一次，求返回能组成target的最小硬币数量，如果不能则返回-1

// 解题：递归
func minCoins(coins []int, target int) int {
	n := len(coins)
	if n == 0 || (n == 1 && coins[0] < target) {
		return -1
	}

	// coins表示硬币数组，index表示当前来到的下标，rest表示选择后剩余金额，返回值表示用了多少枚硬币
	var process func(index, rest int) int

	process = func(index, rest int) int {
		if rest < 0 {
			return -1
		}
		if rest == 0 {
			return 0
		}
		//	rest > 0时
		if index == n {
			return -1
		}
		// 过滤-1的干扰
		p1 := process(index+1, rest)
		p2 := process(index+1, rest-coins[index]) // 选择，后续要+1
		if p1 == -1 && p2 == -1 {
			return -1
		} else {
			if p1 == -1 {
				return p2 + 1
			}
			if p2 == -1 {
				return p1
			}
		}

		return min(p1, p2+1)
	}

	return process(0, target)
}

// 解题：动态规划 => 本质是数组中去拼凑target的整数，
// 于是，构建dp数组，二维的，dp[index][target] :表示选择或者不选择index的硬币，解决target的钱
// dp[0][target] = min(dp[0][target], dp[0][target-coins[0]]
func minCoinsWithDP(coins []int, target int) int {
	n := len(coins)
	if n == 0 || (n == 1 && coins[0] < target) {
		return -1
	}

	dp := make([][]int, n+1) // dp[index][target]
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 0
	}
	for i := 1; i <= target; i++ {
		dp[n][i] = -1
	}

	// 填表，从下往上填写
	for index := n - 1; index >= 0; index-- {
		for rest := 1; rest <= target; rest++ {
			p1 := dp[index+1][rest]
			p2 := -1
			if rest-coins[index] >= 0 {
				p2 = dp[index+1][rest-coins[index]]
			}
			if p1 == -1 && p2 == -1 {
				dp[index][rest] = -1
			} else {
				if p1 == -1 {
					dp[index][rest] = p2 + 1
				} else if p2 == -1 {
					dp[index][rest] = p1
				} else {
					dp[index][rest] = min(p1, p2+1)
				}
			}
		}
	}

	return dp[0][target]
}

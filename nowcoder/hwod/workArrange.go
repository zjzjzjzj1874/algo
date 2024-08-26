package hwod

// 工作安排
// 题目描述
// 小明每周上班都会拿到自己的工作清单，工作清单内包含 n 项工作，每项工作都有对应的 耗时时长(单位 h)和报酬，工作的总报酬为所有已完成工作的报酬之和。那么请你帮小 明安排一下工作，保证小明在指定的工作时间内工作收入最大化。
// 输入描述
// 输入的第一行为两个正整数 T，n。T 代表工作时长(单位 h，0 < T < 100000)，n 代表工作数量(1 < n ≤ 3000)。
// 接下来是 n 行，每行包含两个整数 t，w。t 代表该项工作消耗的时长(单位 h，t > 0)，w 代表该项工作的报酬。
// 输出描述 输出小明指定工作时长内工作可获得的最大报酬。
func maxProfitWithDP(T int, tasks [][]int) int {
	dp := make([]int, T+1)

	for _, task := range tasks {
		time := task[0]
		reward := task[1]

		for j := T; j >= time; j-- {
			if dp[j-time]+reward > dp[j] {
				dp[j] = dp[j-time] + reward
			}
		}
	}
	return dp[T]
}

func maxProfit(T int, tasks [][]int) int {
	dp := make([]int, T+1)

	for _, task := range tasks {
		time := task[0]
		reward := task[1]
		// 从后向前更新
		for j := T; j >= time; j-- {
			if dp[j-time]+reward > dp[j] {
				dp[j] = dp[j-time] + reward
			}
		}
	}

	return dp[T]
}

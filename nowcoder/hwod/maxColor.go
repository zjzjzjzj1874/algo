package hwod

// 找出通过车辆最多颜色
// 题目描述
// 在一个狭小的路口，每秒只能通过一辆车，假如车辆的颜色只有 3 种，找出 N 秒内经过的最 多颜色的车辆数量
// 三种颜色编号为 0,1,2
// 输入描述
// 第一行输入的是通过的车辆颜色信息
// [0,1,1,2] 代表 4 秒钟通过的车辆颜色分别是 0,1,1,2 第二行输入的是统计时间窗，整型，单位为秒
// 输出描述 输出指定时间窗内经过的最多颜色的车辆数量
func maxColor(n int, cars []int) (ans int) {
	// 滑动窗口
	carMap := make(map[int]int)
	l := 0
	for r := 0; r < len(cars); r++ {
		carMap[cars[r]]++
		if r-l+1 > n {
			carMap[cars[l]]--
			l++
		}
		ans = max(ans, carMap[0], carMap[1], carMap[2])
	}

	return
}

func maxProfitV1(T int, tasks [][]int) int {
	dp := make([]int, T+1)

	for _, task := range tasks {
		time := task[0]
		reward := task[1]

		for i := T; i >= time; i-- {
			if dp[i-time]+reward > dp[i] {
				dp[i] = dp[i-time] + reward
			}
		}
	}

	return dp[T]
}

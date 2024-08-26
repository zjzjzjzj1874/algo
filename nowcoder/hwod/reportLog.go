package hwod

// 日志采集系统
// 题目描述
// 日志采集是运维系统的的核心组件。日志是按行生成，每行记做一条，由采集系统分批上 报。 如果上报太频繁，会对服务端造成压力;如果上报太晚，会降低用户的体验;如果一次上 报的条数太多，会导致超时失败。
// 为此，项目组设计了如下的上报策略: 1、每成功上报一条日志，奖励 1 分
// 2、每条日志每延迟上报 1 秒，扣 1 分
// 3、积累日志达到 100 条，必须立即上报 给出日志序列，根据该规则，计算首次上报能获得的最多积分数
// 输入描述
// 按时序产生的日志条数 T1,T2...Tn, 其中 1<=n<=1000,0<=Ti<=100
// 输出描述 首次上报最多能获得的积分数

// 解题
func reportLog(logs []int) (ans int) {
	n := len(logs)
	if n == 1 {
		return n
	}

	total := logs[0]
	ans = logs[0]
	for i := 1; i < n && total <= 100; i++ {
		total += logs[i]
		delay := 0
		for j := i - 1; j >= 0; j-- {
			delay += (i - j) * logs[j]
		}
		ans = max(ans, total-delay)
	}

	return
}

func reportLogWithGPT(logs []int) (ans int) {
	maxScore := 0
	time := 0
	currentScore := 0

	for _, task := range logs {
		time += 1
		currentScore += 1

		if time > 0 {
			currentScore -= time
		}

		maxScore = max(maxScore, currentScore)

		if task == 100 {
			currentScore = 100
			break
		}
	}
	return maxScore
}

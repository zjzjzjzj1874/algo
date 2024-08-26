package hwod

import "sort"

// 任务混部
// 题目描述
// 公司创新实验室正在研究如何最小化资源成本，最大化资源利用率，请你设计算法帮他们 解决一个任务混部问题:有 taskNum 项任务，每个任务有开始时间(startTime )，结束时间(endTime)，并行度(parallelism)三个属性，并行度是指这个任务运行时将 会占用的服务器数量，一个服务器在每个时刻可以被任意任务使用但最多被一个任务占用 ，任务运行完会立即释放(结束时刻不占用)。任务混部问题是指给定一批任务，让这批 任务由同一批服务器承载运行，请你计算完成这批任务混部最少需要多少服务器，从而最 大化控制资源成本。
// 输入描述
// 第一行输入为 taskNum，表示有 taskNum 项任务
// 接下来 taskNum 行，每行三个整数，表示每个任务的开始时间(startTime )，结束时间(endTime)，并行度(parallelism)
// 输出描述 一个整数，表示最少需要的服务器数量

// 解题：和合并区间类似，把重合的放在一起
// [[1,3,3],[2,4,1],[4,7,1],[5,6,2]]

// 1 2 3 4 5 6 7
// 1 - 3
//
//	2 - 4
//	    4  - - 7
//	      5 - 6
func serverNum(taskNum int, tasks [][]int) (ans int) {
	events := make([][2]int, 0, len(tasks)*2)

	// 构造事件点，开始时间增加服务器需求，结束时间减少服务器需求
	for _, task := range tasks {
		startTime, endTime, parallelism := task[0], task[1], task[2]
		events = append(events, [2]int{startTime, parallelism})
		events = append(events, [2]int{endTime, -parallelism})
	}

	// 按时间顺序排序事件点，如果时间相同，则先处理减少服务器需求的事件
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	currentServers := 0

	// 遍历事件点，计算并更新所需的最大服务器数量
	for _, event := range events {
		currentServers += event[1]
		if currentServers > ans {
			ans = currentServers
		}
	}

	return
}

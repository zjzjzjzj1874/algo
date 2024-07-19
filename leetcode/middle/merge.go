package middle

import (
	"sort"
)

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 解题：先对数组的首个元素排序，然后相邻的元素决定是否合并，有个中间数据存储
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals)/2)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] >= intervals[i][0] { // 有重合，需要prev继续与下一个判断
			if prev[1] < intervals[i][1] {
				prev[1] = intervals[i][1]
			}
		} else { // 没有重合
			res = append(res, prev)
			prev = intervals[i] // 加入后prev更新
		}
	}

	res = append(res, prev)
	return res
}

func mergeWithNew(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	// 按首元素从小打到对数组排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0, len(intervals))
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}

	ans = append(ans, []int{start, end})

	return ans
}

// 有这样一个题：一个二维数组，这样[[1,3],[2,6],[8,10],[15,18]],表示会议的开始和结束时间，每个主持人可以主持的会议数量不限制，但是时间不能重复，请算出至少需要的主持人数量？
//[[1,3],[2,6],[8,10],[15,18]] => 2
//[[1,4],[2,5],[3,6],[4,7]] => 3
//[[1,6],[2,5],[3,7],[4,7]] => 4     [2,5][1,6][3,7][4,7]
// 解题：这个题本质应该是合并区间，然后返回ans的len。合并规则：前面结束时间要小于后面的开始时间

func meetingHolder(meets [][]int) int {
	if len(meets) < 2 {
		return len(meets)
	}
	n := len(meets)
	com := make(map[int]bool) // idx是否合并过
	ans := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		if com[i] { // 已经组合过
			continue
		}
		start := meets[i][0]
		end := meets[i][1]
		for j := 1; j < n; j++ {
			if end <= meets[j][0] {
				com[j] = true
				if end < meets[j][1] {
					end = meets[j][1]
				}
			}
		}

		com[i] = true
		ans = append(ans, []int{start, end})
	}

	return len(ans)
}

// 252.会议室问题II
func meetingRooms(meets [][]int) int {
	if len(meets) == 0 {
		return 0
	}

	// 提取所有的开始时间和结束时间
	starts := make([]int, len(meets))
	ends := make([]int, len(meets))
	for i, interval := range meets {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	// 对开始时间和结束时间进行排序
	sort.Ints(starts)
	sort.Ints(ends)

	startPointer, endPointer := 0, 0
	roomsNeeded := 0

	// 遍历所有的开始时间
	for startPointer < len(meets) {
		if starts[startPointer] < ends[endPointer] {
			roomsNeeded++
			startPointer++
		} else {
			endPointer++
			startPointer++
		}
	}

	return roomsNeeded
}

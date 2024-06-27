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

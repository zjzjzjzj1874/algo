package middle

import "sort"

// 1029. 两地调度
// 公司计划面试 2n 人。给你一个数组 costs ，其中 costs[i] = [aCosti, bCosti] 。第 i 人飞往 a 市的费用为 aCosti ，飞往 b 市的费用为 bCosti 。
//
// 返回将每个人都飞到 a 、b 中某座城市的最低费用，要求每个城市都有 n 人抵达。
//
// 示例 1：
//
// 输入：costs = [[10,20],[30,200],[400,50],[30,20]]
// 输出：110
// 解释：
// 第一个人去 a 市，费用为 10。
// 第二个人去 a 市，费用为 30。
// 第三个人去 b 市，费用为 50。
// 第四个人去 b 市，费用为 20。
//
// 最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。
// 示例 2：
//
// 输入：costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
// 输出：1859
// 示例 3：
//
// 输入：costs = [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]
// 输出：3086
//
// 提示：
//
// 2 * n == costs.length
// 2 <= costs.length <= 100
// costs.length 为偶数
// 1 <= aCosti, bCosti <= 1000
func twoCitySchedCost(costs [][]int) int {
	// 解题：计算这两市的差值，然后排序
	n := len(costs)
	diffMap := make(map[int][]int) // map[diffVal][]indexs
	diffVal := make([]int, 0, n)
	for i, cost := range costs {
		diff := cost[1] - cost[0] // 表示B市和A市的费用差值
		diffVal = append(diffVal, diff)
		_, ok := diffMap[diff]
		if !ok {
			diffMap[diff] = []int{i}
		} else {
			diffMap[diff] = append(diffMap[diff], i)
		}
	}

	sort.Ints(diffVal)
	// 排序后，值越小应该去B市，值越大应该去A市, 值相等表示AB两市均可
	ans := 0
	for i := 0; i < n/2; i++ {
		diffB := diffVal[i]     // 选择第二个
		diffA := diffVal[n-1-i] // 选择第一个
		if diffA == diffB {
			idxs := diffMap[diffA]
			// 至少有两个
			a := idxs[len(idxs)-1]
			b := idxs[len(idxs)-2]
			ans += costs[a][0] + costs[b][1]
			idxs = idxs[0 : len(idxs)-2]
			if len(idxs) > 0 {
				diffMap[diffA] = idxs
			}
		} else {
			ia := diffMap[diffA]
			ib := diffMap[diffB]
			// 至少有一个
			a := ia[len(ia)-1]
			b := ib[len(ib)-1]
			ans += costs[a][0] + costs[b][1]
			ia = ia[0 : len(ia)-1]
			ib = ib[0 : len(ib)-1]
			if len(ia) > 0 {
				diffMap[diffA] = ia
			}
			if len(ib) > 0 {
				diffMap[diffB] = ib
			}
		}
	}

	return ans
}

// 解题：减少空间的使用
func twoCitySchedCostWithO1(costs [][]int) (ans int) {
	// 解题：计算这两市的差值，然后排序
	sort.Slice(costs, func(i, j int) bool {
		// 排序,按照1-0元素的差值从小打到排序
		return costs[i][1]-costs[i][0] < costs[j][1]-costs[j][0]
	})

	// 排序后，值越小应该去B市，值越大应该去A市, 值相等表示AB两市均可
	n := len(costs)
	for i := 0; i < n/2; i++ {
		ans += costs[i][1] + costs[n-1-i][0]
	}

	return ans
}

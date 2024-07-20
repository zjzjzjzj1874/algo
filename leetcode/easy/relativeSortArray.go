package easy

import "sort"

// LCR 075. 数组的相对排序
// 给定两个数组，arr1 和 arr2，
//
// arr2 中的元素各不相同
// arr2 中的每个元素都出现在 arr1 中
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
//
// 示例：
//
// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// 输出：[2,2,2,1,4,3,3,9,6,7,19]
//
// 提示：
//
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// arr2 中的元素 arr2[i] 各不相同
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中
//
// 注意：本题与主站 1122 题相同：https://leetcode-cn.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	aMap := make(map[int]int) // map[arr1]count
	for i := range arr1 {
		aMap[arr1[i]]++
	}
	ans := make([]int, 0, len(arr1)+len(arr2))
	helper := make([]int, 0, len(arr1)-len(arr2))

	for i := 0; i < len(arr2); i++ {
		ans = append(ans, arr2[i])
		val, ok := aMap[arr2[i]]
		if !ok {
			continue
		}
		for j := 0; j < val-1; j++ {
			ans = append(ans, arr2[i])
		}
		delete(aMap, arr2[i])
	}

	for key, val := range aMap {
		for j := 0; j < val; j++ {
			helper = append(helper, key)
		}
	}
	sort.Ints(helper)
	ans = append(ans, helper...)

	return ans
}

func relativeSortArrayWithLT(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	rank := make(map[int]int) // map[arr2]index
	for i := range arr2 {
		rank[arr2[i]] = i
	}
	sort.Slice(arr1, func(i int, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			// 如果 hasX 为 true 而 hasY 为 false，return hasX 将返回 true，表示 x 排在 y 之前。
			// 如果 hasX 为 false 而 hasY 为 true，return hasX 将返回 false，表示 x 排在 y 之后。
			return hasX
		}
		return x < y
	})

	return arr1
}

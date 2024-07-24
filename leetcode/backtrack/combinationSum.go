package backtrack

import "sort"

// LCR 081. 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
//
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
//
// 示例 1：
//
// 输入: candidates = [2,3,6,7], target = 7
// 输出: [[7],[2,2,3]]
// 示例 2：
//
// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]
// 示例 3：
//
// 输入: candidates = [2], target = 1
// 输出: []
// 示例 4：
//
// 输入: candidates = [1], target = 1
// 输出: [[1]]
// 示例 5：
//
// 输入: candidates = [1], target = 2
// 输出: [[1,1]]
//
// 提示：
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都是独一无二的。
// 1 <= target <= 500
//
// 注意：本题与主站 39 题相同： https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) (ans [][]int) {
	selected := make([]int, 0, 4)
	selectSum := 0

	var process func(start int)

	process = func(start int) {
		if selectSum == target {
			res := append([]int{}, selected...)
			ans = append(ans, res)
			return
		}

		// 从start开始，避免重复统计
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			// 修剪枝叶，当sum>target，无需进入递归
			if selectSum+num > target {
				continue
			}
			selected = append(selected, num)
			selectSum += num
			process(i)

			selectSum -= num
			selected = selected[:len(selected)-1]
		}
	}

	process(0)

	return
}

// LCR 082. 组合总和 II
// 给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
//
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
//
// 提示:
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// 注意：本题与主站 40 题相同： https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sum := 0
	choice := make([]int, 0, 4)
	sort.Ints(candidates) // 排序

	var backtrack func(start int)
	backtrack = func(start int) {
		if sum == target {
			res := append([]int{}, choice...)
			ans = append(ans, res)
			return
		}

		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			if sum+num > target { // 和超过目标值 || 被选择过
				continue
			}
			if i > start && num == candidates[i-1] {
				continue
			}

			sum += num
			choice = append(choice, num)
			backtrack(i + 1) // 不能重复选择，就直接选择下一个

			// 回溯 = 撤销
			sum -= num
			choice = choice[:len(choice)-1]
		}
	}

	backtrack(0)

	return
}

package backtrack

import "fmt"

// LCR 083. 全排列
// 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
//
// 输入：nums = [1]
// 输出：[[1]]
//
// 提示：
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// 注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/

// 解题：回溯
func permuteWithDfs(nums []int) (ans [][]int) {
	n := len(nums)
	choice := make([]int, 0, n)
	selected := make([]bool, n)

	var dfs func()
	dfs = func() {
		if n == len(choice) { // 拼好一个答案
			ans = append(ans, append([]int{}, choice...))
			return
		}

		for i := 0; i < n; i++ {
			if selected[i] {
				continue
			}

			selected[i] = true
			choice = append(choice, nums[i])
			dfs()

			selected[i] = false
			choice = choice[:len(choice)-1]
		}
	}

	dfs()
	return
}

// 解题：回溯
func permute(nums []int) (ans [][]int) {
	choice := make([]int, 0, len(nums))
	selected := make([]bool, len(nums))

	process(&nums, &choice, &selected, &ans)
	return
}

func process(nums, choice *[]int, selected *[]bool, res *[][]int) {
	if len(*nums) == len(*choice) {
		newState := append([]int{}, *choice...) // 这里不能直接*res = append(*res, *choice) // 因为存的是地址，如果这样写，获取的结果都一样了
		*res = append(*res, newState)           // 记录结果
	}

	// 遍历数组
	for i := range *nums {
		if (*selected)[i] { // 之前被选择过，直接跳过
			continue
		}

		*choice = append(*choice, (*nums)[i]) // 选择当前元素
		(*selected)[i] = true                 // 变更选择状态
		fmt.Println("递归之前：", choice)
		fmt.Println("递归之前：", selected)
		process(nums, choice, selected, res)
		(*selected)[i] = false // 撤销选择
		*choice = (*choice)[:len(*choice)-1]
		fmt.Println("	递归之后：", choice)
		fmt.Println("	递归之后：", selected)
	}
}

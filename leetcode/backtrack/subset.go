package backtrack

// LCR 079. 子集
// 给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：
//
// 输入：nums = [0]
// 输出：[[],[0]]
//
// 提示：
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// 注意：本题与主站 78 题相同： https://leetcode-cn.com/problems/subsets/

// 解题：子集，子集包含空子集，单独处理，后面遍历nums的长度，然后固定子集的长度
func subsets(nums []int) (ans [][]int) {
	var backtrack func(start, k int, choice *[]int)
	backtrack = func(start, k int, choice *[]int) {
		if len(*choice) == k {
			res := append([]int{}, (*choice)...)
			ans = append(ans, res)
			return
		}

		for i := start; i < len(nums); i++ {
			*choice = append(*choice, nums[i])

			backtrack(i+1, k, choice)

			*choice = (*choice)[:len(*choice)-1]
		}
	}

	ans = append(ans, []int{})
	for i := 0; i < len(nums); i++ {
		choice := make([]int, 0, i)
		backtrack(0, i+1, &choice)
	}

	return
}

// 解题：深度优先遍历:输入的视角，可以选择也可以不选择
func subsetsWithDFSInput(nums []int) (ans [][]int) {
	n := len(nums)
	choice := make([]int, 0, n)
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			ans = append(ans, append([]int{}, choice...))
			return
		}

		// 不选择[start]
		dfs(start + 1)

		// 选择nums[start]
		choice = append(choice, nums[start])
		dfs(start + 1)
		choice = choice[:len(choice)-1] // 恢复之前的状态
	}

	dfs(0)
	return
}

// 解题：深度优先遍历:答案的视角，即第一个选谁，第二个选谁，第三个选谁，当然每个都有选择或者不选择的选项，即2^n个选项
func subsetsWithDFSAns(nums []int) (ans [][]int) {
	n := len(nums)
	choice := make([]int, 0, n)

	var dfs func(selectIdx int)
	dfs = func(selectIdx int) {
		ans = append(ans, append([]int{}, choice...))

		for i := selectIdx; i < n; i++ {
			choice = append(choice, nums[i]) // 选择
			dfs(i + 1)
			choice = choice[:len(choice)-1] // 恢复选择之前
		}
	}

	dfs(0)
	return
}

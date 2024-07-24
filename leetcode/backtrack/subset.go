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

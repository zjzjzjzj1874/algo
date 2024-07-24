package backtrack

import "fmt"

// LCR 084. 全排列 II
// 给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
//
// 示例 1：
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
// [1,2,1],
// [2,1,1]]
// 示例 2：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 提示：
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// 注意：本题与主站 47 题相同： https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) (ans [][]int) {
	choice := make([]int, 0)            // 已经选择的元素，等待加入ans结果集
	selected := make([]bool, len(nums)) // Index对应的元素是否被选择

	var process func()

	process = func() {
		if len(choice) == len(nums) {
			res := append([]int{}, choice...)
			ans = append(ans, res)
		}

		deplicateMap := make(map[int]struct{}) // 去重
		for i, num := range nums {
			// 有重复元素或者被选择过，直接跳过 =》剪枝
			if _, ok := deplicateMap[num]; ok || selected[i] {
				continue
			}

			choice = append(choice, nums[i])
			deplicateMap[nums[i]] = struct{}{}
			selected[i] = true
			fmt.Println("递归之前：", choice)
			fmt.Println("递归之前：", selected)
			process()
			// 回溯
			selected[i] = false
			choice = choice[:len(choice)-1]
			fmt.Println("	递归之后：", choice)
			fmt.Println("	递归之后：", selected)
		}
	}

	process()

	return ans
}

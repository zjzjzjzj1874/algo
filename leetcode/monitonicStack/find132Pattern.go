package monitonicStack

import "math"

// 456. 132 模式
// 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
//
// 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4]
// 输出：false
// 解释：序列中不存在 132 模式的子序列。
// 示例 2：
//
// 输入：nums = [3,1,4,2]
// 输出：true
// 解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
// 示例 3：
//
// 输入：nums = [-1,3,2,0]
// 输出：true
// 解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
//
// 提示：
//
// n == nums.length
// 1 <= n <= 2 * 105
// -109 <= nums[i] <= 109

// 单调栈
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// 单调栈，大的先入栈， 大 --> 小
	stack := make([]int, 0, len(nums))
	second := math.MinInt // 初始化为非常小的值，表示132模式中的“2”
	for i := n - 1; i >= 0; i-- {
		if nums[i] < second {
			return true
		}

		// 存在这种模式，栈不为空，当前元素比栈顶元素大，则要弹出栈顶元素
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			if stack[len(stack)-1] > second {
				second = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1] // 移除栈顶元素
		}

		stack = append(stack, nums[i])
	}

	return false
}

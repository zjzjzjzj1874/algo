package middle

// 260. 只出现一次的数字 III
// 给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
// 你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
//
// 示例 1：
//
// 输入：nums = [1,2,1,3,2,5]
// 输出：[3,5]
// 解释：[5, 3] 也是有效的答案。
// 示例 2：
//
// 输入：nums = [-1,0]
// 输出：[-1,0]
// 示例 3：
//
// 输入：nums = [0,1]
// 输出：[1,0]
//
// 提示：
//
// 2 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次

// 解题：位运算
func singleNumber(nums []int) []int {
	ab := 0

	for _, num := range nums {
		// 所有的数做异或运算
		ab ^= num
	}
	// ab = a^b

	// a!=b,则可以找出他们最右边的1
	rightOne := ab & (-ab)

	a := 0
	b := 0
	for i := range nums {
		if (rightOne & nums[i]) > 0 { // 这和A是一类数
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}

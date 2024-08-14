package prefixsum

// 3152. 特殊数组 II
// 如果数组的每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个 特殊数组 。
//
// 你有一个整数数组 nums 和一个二维整数矩阵 queries，对于 queries[i] = [fromi, toi]，请你帮助你检查
// 子数组
// nums[fromi..toi] 是不是一个 特殊数组 。
//
// 返回布尔数组 answer，如果 nums[fromi..toi] 是特殊数组，则 answer[i] 为 true ，否则，answer[i] 为 false 。
//
// 示例 1：
//
// 输入：nums = [3,4,1,2,6], queries = [[0,4]]
//
// 输出：[false]
//
// 解释：
//
// 子数组是 [3,4,1,2,6]。2 和 6 都是偶数。
//
// 示例 2：
//
// 输入：nums = [4,3,1,6], queries = [[0,2],[2,3]]
//
// 输出：[false,true]
//
// 解释：
//
// 子数组是 [4,3,1]。3 和 1 都是奇数。因此这个查询的答案是 false。
// 子数组是 [1,6]。只有一对：(1,6)，且包含了奇偶性不同的数字。因此这个查询的答案是 true。
//
// 提示：
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105
// 1 <= queries.length <= 105
// queries[i].length == 2
// 0 <= queries[i][0] <= queries[i][1] <= nums.length - 1

// 解题：怎么快速知道是否有奇偶性相同的相邻元素？
// 考虑这样一个问题：给你一个只包含 0 和 1 的数组，如何快速判断一个子数组是否全为 0？
// 解答：如果子数组的元素和等于 0，那么子数组一定全为 0；如果子数组的元素和大于 0，那么子数组一定包含 1。如何快速计算子数组元素和？这可以用 前缀和 解决。
// 所以我们让相邻元素的奇偶性如果不同，两个元素都相等，则s[i] = s[i-1]; 如果奇偶性相同，那么s[i] = s[i-1]+1，最后判断end和begin的值是否相等即可。
func isArraySpecial(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums)) // 前缀和预处理元素
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] // 默认两数奇偶性不同，那么s[i] = s[i-1]
		// 如果 nums[i-1] % 2 == nums[i] % 2，则表示这两个元素的奇偶性相同，
		// 因此 s[i] = s[i-1] + 1
		if nums[i-1]%2 == nums[i]%2 {
			s[i]++
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}

	return ans
}

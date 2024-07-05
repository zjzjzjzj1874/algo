package middle

// 442. 数组中重复的数据
// 给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
// 示例 1：
// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[2,3]
// 示例 2：
// 输入：nums = [1,1,2]
// 输出：[1]
// 示例 3：
// 输入：nums = [1]
// 输出：[]
// 提示：
// n == nums.length
// 1 <= n <= 105
// 1 <= nums[i] <= n
// nums 中的每个元素出现 一次 或 两次

// 解题：nums的数范围在1，n，出现的次数为1，2次，原地hash，在对应位置上不存在的数字，便是多出现的数；
func findDuplicates(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] != nums[nums[i]-1] {
			// 举例：如果i=0 nums = []int{2,3,5,4} ，则nums[0] = 2,nums[nums[i]-1]表示2该出现的位置不是2，则进行交换
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	ans := make([]int, 0, 4)
	for i, v := range nums {
		if v != i+1 {
			ans = append(ans, v)
		}
	}

	return ans
}

package easy

// 448. 找到所有数组中消失的数字
// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
// 示例 1：
// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[5,6]
// 示例 2：
// 输入：nums = [1,1]
// 输出：[2]
// 提示：
// n == nums.length
// 1 <= n <= 105
// 1 <= nums[i] <= n
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。

// 解题：使用原地hash，因为nums[i] 在区间 [1, n] 内，所以直接按照hash对数组排序。把数组按照顺序一次放入数组中；
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] != nums[nums[i]-1] { // 如果这两个对应的位置不相等，需要进行交换
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, 4)
	for i, v := range nums {
		if v != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}

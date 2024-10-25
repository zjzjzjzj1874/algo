package greedy

// 910. 最小差值 II
// 给你一个整数数组 nums，和一个整数 k 。
// 对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。
// nums 的 分数 是 nums 中最大元素和最小元素的差值。
// 在更改每个下标对应的值之后，返回 nums 的最小 分数 。
// 示例 1：
// 输入：nums = [1], k = 0
// 输出：0
// 解释：分数 = max(nums) - min(nums) = 1 - 1 = 0 。
// 示例 2：
// 输入：nums = [0,10], k = 2
// 输出：6
// 解释：将数组变为 [2, 8] 。分数 = max(nums) - min(nums) = 8 - 2 = 6 。
// 示例 3：
// 输入：nums = [1,3,6], k = 3
// 输出：3
// 解释：将数组变为 [4, 6, 3] 。分数 = max(nums) - min(nums) = 6 - 3 = 3 。
// 提示：
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 104
// 0 <= k <= 104
// 解题：贪心，TODO
func smallestRangeII(nums []int, k int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	mid := 0
	mai := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[mai] {
			mai = i
		}
		if nums[i] < nums[mid] {
			mid = i
		}
	}

	nums[mid] += k
	nums[mai] -= k
	nmi := sliceMin(nums)
	nma := sliceMax(nums)

	if nmi > nums[mid] {
		nmi = nums[mid]
	}
	if nma > nums[mai] {
		nma = nums[mai]
	}

	return nma - nmi
}

func sliceMin(nums []int) int {
	mi := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mi {
			mi = nums[i]
		}
	}

	return mi
}
func sliceMax(nums []int) int {
	ma := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ma {
			ma = nums[i]
		}
	}

	return ma
}

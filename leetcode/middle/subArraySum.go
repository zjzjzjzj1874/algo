package middle

// 560. 和为 K 的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2
// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2
// 提示：
// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

// 枚举，使用双重循环
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			if sum+nums[j] == k {
				count++
			}
			sum += nums[j]
		}
	}

	return count
}

// 前缀和+hash
// 解题：和为K的子序列，保证sum[i] = sum[i-1] + nums[i];则 ==> sum[i] - sum[j-1] = k ==> sum[j-1] = sum[i]-k;
// 使用一个hash表来存储前N项之和，key为SUM，value为SUM出现次数；
func subarraySumV1(nums []int, k int) int {
	count := 0
	preSum := 0
	cacheSumMap := make(map[int]int)

	cacheSumMap[0]++

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := cacheSumMap[preSum-k]; ok {
			count += cacheSumMap[preSum-k]
		}

		cacheSumMap[preSum]++
	}

	return count
}

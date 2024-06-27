package hard

import "math"

// 41. 缺失的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
// 示例 1：
// 输入：nums = [1,2,0]
// 输出：3
// 解释：范围 [1,2] 中的数字都在数组中。
// 示例 2：
// 输入：nums = [3,4,-1,1]
// 输出：2
// 解释：1 在数组中，但 2 没有。
// 示例 3：
// 输入：nums = [7,8,9,11,12]
// 输出：1
// 解释：最小的正数 1 没有出现。
// 提示：
// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1

// 解题：先遍历一遍数组，找出其中最大最小的，并把最大最小之间的数存到hash表中；第二次遍历最大最小之间的数，去hash表中找，不存在就返回
func firstMissingPositive(nums []int) int {
	min, max := math.MaxInt, 0
	exist := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 {
			continue
		}

		exist[nums[i]] = struct{}{}
		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}
	if min > 1 {
		return 1
	}

	for i := min + 1; i < max; i++ {
		if _, ok := exist[i]; !ok {
			return i
		}
	}

	return max + 1
}

// 解题：原地hash；长度为N，那么如果是连续的正整数，其值一定是1，n之间的；这个时候第一次循环，把下表i放置的元素值为i+1；
// 第二次循环，不在自己位置的元素一定是缺失的这个正整数，否则返回n+1；
func firstMissingPositiveV1(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1] // 交换位置
		}
	}

	for i, x := range nums {
		if x != i+1 {
			return i + 1
		}
	}

	return n + 1
}

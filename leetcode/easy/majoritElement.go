package easy

// Leetcode-169: 多数元素 => https://leetcode.cn/problems/majority-element/description/
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 示例 1：
// 输入：nums = [3,2,3]
// 输出：3
// 示例 2：
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
// 提示：
// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

// leetcode的解题思路：1. 相消法，任意两个不相同的元素，同时消失，最后剩余的最多的元素一定是满足题意的。 2. 哈希值法
func majoritElement(nums []int) int {
	res := nums[0]
	majorTimes := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			majorTimes++
		} else {
			if majorTimes == 0 {
				res = nums[i]
				majorTimes++
			} else {
				majorTimes--
			}
		}
	}
	return res
}

func majoritElementWithMore(nums []int) int {
	ele := 0
	vote := 0
	for i := range nums {
		if ele == 0 && vote == 0 { // 第一次来，没有元素也没有投票信息
			ele = nums[i]
			vote++
		} else if ele != 0 && vote == 0 { // 投票被相消了，但是ele是消除的元素
			vote++
			ele = nums[i]
		} else if nums[i] == ele { // 不相同且投票信息不为0
			vote++
		} else {
			vote--
		}
	}

	check := 0
	for i := range nums {
		if nums[i] == ele {
			check++
		}
	}

	if check > len(nums)/2 && vote > 0 {
		return ele
	} else {
		return -1
	}
}
func majoritElementWithMoreV1(nums []int) int {
	ele := 0
	vote := 0
	for i := range nums {
		if vote == 0 {
			ele = nums[i]
		}
		if nums[i] == ele {
			vote++
		} else {
			vote--
		}
	}

	return ele
}

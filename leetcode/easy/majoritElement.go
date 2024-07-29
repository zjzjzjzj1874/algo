package easy

// 169: 多数元素 => https://leetcode.cn/problems/majority-element/description/
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

// 229. 多数元素 II
// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
// 示例 1：
//
// 输入：nums = [3,2,3]
// 输出：[3]
// 示例 2：
//
// 输入：nums = [1]
// 输出：[1]
// 示例 3：
//
// 输入：nums = [1,2]
// 输出：[1,2]
//
// 提示：
//
// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

// 解题：hash表记录
func majorityElement(nums []int) []int {
	n := len(nums)
	least := n / 3 // 至少要大于least

	nMap := make(map[int]int)

	for idx := range nums {
		nMap[nums[idx]]++
	}

	ans := make([]int, 0, len(nMap))
	for key, val := range nMap {
		if val > least {
			ans = append(ans, key)
		}
	}

	return ans
}

func majorityElementWithMor(nums []int) []int {
	n := len(nums)
	ele1, ele2 := 0, 0
	vote1, vote2 := 0, 0

	for i := range nums {
		if vote1 > 0 && nums[i] == ele1 { // 元素1，vote1++
			vote1++
		} else if vote2 > 0 && nums[i] == ele2 { // 元素2，vote2++
			vote2++
		} else if vote1 == 0 { // 第一个元素
			ele1 = nums[i]
			vote1++
		} else if vote2 == 0 { // 第二个元素
			ele2 = nums[i]
			vote2++
		} else { // 出现的三个元素不相同，互相抵消一次
			vote1--
			vote2--
		}
	}

	// 至此，算出来出现频率最高的ele1，ele2
	// 验证一下是否都出现至少n/3
	check1, check2 := 0, 0
	for i := range nums {
		if nums[i] == ele1 {
			check1++
		}
		if nums[i] == ele2 {
			check2++
		}
	}

	ans := make([]int, 0, 2)
	if vote1 > 0 && check1 > n/3 {
		ans = append(ans, ele1)
	}
	if vote2 > 0 && check2 > n/3 {
		ans = append(ans, ele2)
	}

	return ans
}

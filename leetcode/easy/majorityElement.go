package easy

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

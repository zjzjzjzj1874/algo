package bit

// 3158. 求出出现两次数字的 XOR 值
// 提示
// 给你一个数组 nums ，数组中的数字 要么 出现一次，要么 出现两次。
// 请你返回数组中所有出现两次数字的按位 XOR 值，如果没有数字出现过两次，返回 0 。
// 示例 1：
// 输入：nums = [1,2,1,3]
// 输出：1
// 解释：
// nums 中唯一出现过两次的数字是 1 。
// 示例 2：
// 输入：nums = [1,2,3]
// 输出：0
// 解释：
// nums 中没有数字出现两次。
// 示例 3：
// 输入：nums = [1,2,2,1]
// 输出：3
// 解释：
// 数字 1 和 2 出现过两次。1 XOR 2 == 3 。
// 提示：
// 1 <= nums.length <= 50
// 1 <= nums[i] <= 50
// nums 中每个数字要么出现过一次，要么出现过两次。
func duplicateNumbersXOR(nums []int) (ans int) {
	// hash
	nm := make(map[int]int)
	for i := range nums {
		nm[nums[i]]++
	}

	for val, count := range nm {
		if count == 2 {
			ans ^= val
		}
	}

	return
}

func duplicateNumbersXORWithArr(nums []int) (ans int) {
	arr := make([]int, 50)
	// 申请50个槽位，因为数字是从1-50的，所以槽位的index代表具体数字，里面的值表示出现的次数
	for i := range nums {
		arr[nums[i]-1]++
	}

	for idx, count := range arr {
		if count == 2 {
			ans ^= idx + 1
		}
	}

	return
}

// 位运算
func duplicateNumbersXORWithBitOp(nums []int) (ans int) {
	t := 0
	for i := range nums {
		tt := t >> nums[i]
		if tt&1 > 0 { // 已经出现过
			ans ^= nums[i]
		} else {
			t |= 1 << nums[i]
		}
	}

	return
}

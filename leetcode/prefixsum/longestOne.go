package prefixsum

// 1004. 最大连续1的个数 III
// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
//
// 示例 1：
//
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
// 示例 2：
//
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length

// 解题：滑动窗口+前缀和
func longestOnes(nums []int, k int) (ans int) {
	l := 0
	lZero := 0
	rZero := 0

	for r, num := range nums {
		rZero += 1 - num // 这里1-num，其实记录的是0的个数，当num==1时，1-1=0，这个时候rZero的值并不会改变；当num=0时，1-0 = 1，表示此时有一个0

		for lZero < rZero-k {
			lZero += 1 - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}

	return
}

// 485. 最大连续 1 的个数
// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
//
// 示例 1：
//
// 输入：nums = [1,1,0,1,1,1]
// 输出：3
// 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
// 示例 2:
//
// 输入：nums = [1,0,1,1,0,1]
// 输出：2
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1.
func findMaxConsecutiveOnes(nums []int) (ans int) {
	// 解题：左右滑窗+前缀和
	l := 0
	sum := 0
	for r, num := range nums {
		sum += 1 - num

		for sum > 0 {
			sum -= 1 - nums[l]
			l++
		}

		ans = max(ans, r-l+1)
	}

	return
}

// 解题：一次遍历，计算
func findMaxConsecutiveOnesWithOne(nums []int) (ans int) {
	sum := 0
	for _, num := range nums {
		if num == 1 {
			sum++
			ans = max(ans, sum)
		} else {
			sum = 0
		}
	}

	return
}

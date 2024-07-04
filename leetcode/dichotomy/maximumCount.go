package dichotomy

// 2529. 正整数和负整数的最大计数
// 给你一个按 非递减顺序 排列的数组 nums ，返回正整数数目和负整数数目中的最大值。
//
// 换句话讲，如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值。
// 注意：0 既不是正整数也不是负整数。
//
// 示例 1：
//
// 输入：nums = [-2,-1,-1,1,2,3]
// 输出：3
// 解释：共有 3 个正整数和 3 个负整数。计数得到的最大值是 3 。
// 示例 2：
//
// 输入：nums = [-3,-2,-1,0,0,1,2]
// 输出：3
// 解释：共有 2 个正整数和 3 个负整数。计数得到的最大值是 3 。
// 示例 3：
//
// 输入：nums = [5,20,66,1314]
// 输出：4
// 解释：共有 4 个正整数和 0 个负整数。计数得到的最大值是 4 。
//
// 提示：
//
// 1 <= nums.length <= 2000
// -2000 <= nums[i] <= 2000
// nums 按 非递减顺序 排列。
//
// 进阶：你可以设计并实现时间复杂度为 O(log(n)) 的算法解决此问题吗？

// 解题：递增。二分法，先求出负数的下标，再求出正数的下标
func maximumCount(nums []int) int {
	n := len(nums)
	l := 0
	r := n
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	negCount := l

	l = 0
	r = n
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	posCount := n - l

	if posCount > negCount {
		return posCount
	}

	return negCount
}

package greedy

import "math"

// 2091. 从数组中移除最大值和最小值
// 给你一个下标从 0 开始的数组 nums ，数组由若干 互不相同 的整数组成。
//
// nums 中有一个值最小的元素和一个值最大的元素。分别称为 最小值 和 最大值 。你的目标是从数组中移除这两个元素。
//
// 一次 删除 操作定义为从数组的 前面 移除一个元素或从数组的 后面 移除一个元素。
//
// 返回将数组中最小值和最大值 都 移除需要的最小删除次数。
//
// 示例 1：
//
// 输入：nums = [2,10,7,5,4,1,8,6]
// 输出：5
// 解释：
// 数组中的最小元素是 nums[5] ，值为 1 。
// 数组中的最大元素是 nums[1] ，值为 10 。
// 将最大值和最小值都移除需要从数组前面移除 2 个元素，从数组后面移除 3 个元素。
// 结果是 2 + 3 = 5 ，这是所有可能情况中的最小删除次数。
// 示例 2：
//
// 输入：nums = [0,-4,19,1,8,-2,-3,5]
// 输出：3
// 解释：
// 数组中的最小元素是 nums[1] ，值为 -4 。
// 数组中的最大元素是 nums[2] ，值为 19 。
// 将最大值和最小值都移除需要从数组前面移除 3 个元素。
// 结果是 3 ，这是所有可能情况中的最小删除次数。
// 示例 3：
//
// 输入：nums = [101]
// 输出：1
// 解释：
// 数组中只有这一个元素，那么它既是数组中的最小值又是数组中的最大值。
// 移除它只需要 1 次删除操作。
//
// 提示：
//
// 1 <= nums.length <= 105
// -105 <= nums[i] <= 105
// nums 中的整数 互不相同
func minimumDeletions(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return n
	}

	ma := math.MinInt
	mai := -1
	mi := math.MaxInt
	mii := -1
	for i, num := range nums {
		if num > ma {
			ma = num
			mai = i
		}
		if num < mi {
			mi = num
			mii = i
		}
	}
	if mii > mai {
		mii, mai = mai, mii // 交换大小位置
	}

	if n-mai < mii+1 {
		// mai离两头近点，先删除mai
		ans = n - mai
		n = n - ans

		// 大的删除不影响mii的
		if n-mii < mii+1 {
			// 后面删除更近
			ans += n - mii
		} else {
			ans += mii + 1 // 前面删更近
		}
	} else {
		// mii离两头近点，先删除mii
		ans = mii + 1
		n = n - ans
		mai -= ans

		// 判断mai从前面还是后面删除最近
		if n-mai < mai+1 {
			// 后面删除更近
			ans += n - mai
		} else {
			ans += mai + 1 // 前面删更近
		}
	}

	return
}

// 解题：简化逻辑
func minimumDeletionsWithOptimize(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return n
	}

	ma := 0
	mi := 0
	for i, num := range nums {
		if num > nums[ma] {
			ma = i
		}
		if num < nums[mi] {
			mi = i
		}
	}

	if mi > ma {
		mi, ma = ma, mi // 交换大小位置, 保证mi < ma
	}
	// [0, 0, 0, mi, 0, 0, ma, 0, 0]
	// ma+1: 表示从做到ma需要删除的数字
	// n-mi: 表示从右到mi需要删除的位置
	// mi+1+n-ma: 表示0-mi和n-ma的和，需要删除的数字
	// 因为mi < ma是严格成立的，上面三种删除方式严格囊括了大小数删除，取其最小值即为答案
	return min(ma+1, n-mi, mi+1+n-ma)
}

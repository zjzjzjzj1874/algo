package easy

// LCR 012. 寻找数组的中心下标
// 给你一个整数数组 nums ，请计算数组的 中心下标 。
//
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
//
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
//
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
//
// 示例 1：
//
// 输入：nums = [1,7,3,6,5,6]
// 输出：3
// 解释：
// 中心下标是 3 。
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
// 示例 2：
//
// 输入：nums = [1, 2, 3]
// 输出：-1
// 解释：
// 数组中不存在满足此条件的中心下标。
// 示例 3：
//
// 输入：nums = [2, 1, -1]
// 输出：0
// 解释：
// 中心下标是 0 。
// 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
// 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
//
// 提示：
//
// 1 <= nums.length <= 104
// -1000 <= nums[i] <= 1000
//
// 注意：本题与主站 724 题相同： https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	n := len(nums)
	prefixSum := make(map[int]int) // map[index]sum
	suffixSum := make(map[int]int) // map[index]sum

	sum := 0
	ss := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		ss += nums[n-1-i]

		prefixSum[i] = sum
		suffixSum[n-i-1] = ss
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			if suffixSum[i+1] == 0 {
				return i
			}
		} else if i > 0 && i < n-1 {
			if prefixSum[i-1] == suffixSum[i+1] {
				return i
			}
		} else { // i = n-1
			if prefixSum[i-1] == 0 {
				return i
			}
		}
	}

	return -1
}

// 后缀和，前缀和不需要缓存
func pivotIndexWithSingleMap(nums []int) int {
	n := len(nums)
	suffixSum := make(map[int]int) // map[index]sum

	ss := 0
	for i := 0; i < n; i++ {
		ss += nums[n-1-i]

		suffixSum[n-i-1] = ss
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if i == 0 {
			if suffixSum[i+1] == 0 {
				return i
			}
		} else if i > 0 && i < n-1 {
			if sum-nums[i] == suffixSum[i+1] {
				return i
			}
		} else { // i = n-1
			if sum-nums[i] == 0 {
				return i
			}
		}
	}

	return -1
}

func pivotIndexWithoutMap(nums []int) int {
	n := len(nums)

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	tmp := 0
	for i, num := range nums {
		if 2*tmp+num == sum {
			return i
		}
		tmp += num
	}

	return -1
}

package middle

import "sort"

// 3132. 找出与数组相加的整数 II
// 给你两个整数数组 nums1 和 nums2。
//
// 从 nums1 中移除两个元素，并且所有其他元素都与变量 x 所表示的整数相加。如果 x 为负数，则表现为元素值的减少。
//
// 执行上述操作后，nums1 和 nums2 相等 。当两个数组中包含相同的整数，并且这些整数出现的频次相同时，两个数组 相等 。
//
// 返回能够实现数组相等的 最小 整数 x 。
//
// 示例 1:
//
// 输入：nums1 = [4,20,16,12,8], nums2 = [14,18,10]
//
// 输出：-2
//
// 解释：
//
// 移除 nums1 中下标为 [0,4] 的两个元素，并且每个元素与 -2 相加后，nums1 变为 [18,14,10] ，与 nums2 相等。
//
// 示例 2:
//
// 输入：nums1 = [3,5,5,3], nums2 = [7,7]
//
// 输出：2
//
// 解释：
//
// 移除 nums1 中下标为 [0,3] 的两个元素，并且每个元素与 2 相加后，nums1 变为 [7,7] ，与 nums2 相等。
//
// 提示：
//
// 3 <= nums1.length <= 200
// nums2.length == nums1.length - 2
// 0 <= nums1[i], nums2[i] <= 1000
// 测试用例以这样的方式生成：存在一个整数 x，nums1 中的每个元素都与 x 相加后，再移除两个元素，nums1 可以与 nums2 相等。
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	n1 := len(nums1)
	n2 := len(nums2)

	ans := 1001 // 初始化

	for i := 0; i < n1; i++ {
		for j := i + 1; j < n1; j++ {
			cop := make([]int, len(nums1))
			copy(cop, nums1)
			tmp := cop[0:i]
			tmp = append(tmp, cop[i+1:j]...)
			tmp = append(tmp, cop[j+1:]...)

			allEqual := true
			diff := nums2[0] - tmp[0]
			for left, right := 0, n2-1; left <= right; {
				if nums2[left]-tmp[left] != diff || diff != nums2[right]-tmp[right] {
					allEqual = false
					break // 不满足条件，直接走一下个
				}
				left++
				right--
			}
			if allEqual && diff < ans {
				ans = diff
			}
		}
	}

	return ans
}

// 更简单的解题
func minimumAddedIntegerWith(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 2; i >= 0; i-- { // 因为排序后，从右向左，即代表从大到小，这样第一次碰到合适的，就可以直接返回
		left := i + 1
		right := 1
		for left < len(nums1) && right < len(nums2) {
			if nums1[left]-nums2[right] == nums1[i]-nums2[0] {
				right++
			}
			left++
		}
		if right == len(nums2) {
			// 第一次碰到合适的，就直接返回
			return nums2[0] - nums1[i]
		}
	}

	return 0
}

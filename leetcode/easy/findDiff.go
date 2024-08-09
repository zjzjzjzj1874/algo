package easy

import "sort"

// 2215. 找出两数组的不同
// 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：
//
// answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
// answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
// 注意：列表中的整数可以按 任意 顺序返回。
//
// 示例 1：
//
// 输入：nums1 = [1,2,3], nums2 = [2,4,6]
// 输出：[[1,3],[4,6]]
// 解释：
// 对于 nums1 ，nums1[1] = 2 出现在 nums2 中下标 0 处，然而 nums1[0] = 1 和 nums1[2] = 3 没有出现在 nums2 中。因此，answer[0] = [1,3]。
// 对于 nums2 ，nums2[0] = 2 出现在 nums1 中下标 1 处，然而 nums2[1] = 4 和 nums2[2] = 6 没有出现在 nums2 中。因此，answer[1] = [4,6]。
// 示例 2：
//
// 输入：nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// 输出：[[3],[]]
// 解释：
// 对于 nums1 ，nums1[2] 和 nums1[3] 没有出现在 nums2 中。由于 nums1[2] == nums1[3] ，二者的值只需要在 answer[0] 中出现一次，故 answer[0] = [3]。
// nums2 中的每个整数都在 nums1 中出现，因此，answer[1] = [] 。
//
// 提示：
//
// 1 <= nums1.length, nums2.length <= 1000
// -1000 <= nums1[i], nums2[i] <= 1000
func findDifference(nums1 []int, nums2 []int) [][]int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	n1 := len(nums1)
	n2 := len(nums2)

	a1Map := make(map[int]struct{})
	a2Map := make(map[int]struct{})
	left, right := 0, 0
	for left, right = 0, 0; left < n1 && right < n2; {
		if nums1[left] < nums2[right] {
			a1Map[nums1[left]] = struct{}{}
			left++
		} else if nums1[left] == nums2[right] {
			left++
		} else {
			right++
		}
	}
	for left < n1 {
		a1Map[nums1[left]] = struct{}{}
		left++
	}
	for left, right = 0, 0; left < n2 && right < n1; {
		if nums2[left] < nums1[right] {
			a2Map[nums2[left]] = struct{}{}
			left++
		} else if nums2[left] == nums1[right] {
			left++
		} else {
			right++
		}
	}
	for left < n2 {
		a2Map[nums2[left]] = struct{}{}
		left++
	}

	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)
	for i := range a1Map {
		ans[0] = append(ans[0], i)
	}
	for i := range a2Map {
		ans[1] = append(ans[1], i)
	}

	return ans
}

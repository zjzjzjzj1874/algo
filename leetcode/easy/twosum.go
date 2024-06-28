package easy

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
// 提示：
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案
// 进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

// 针对有序数组的双指针
func twosumTwoPtr(arr []int, target int) []int {
	//sort.Ints(arr) // golang中排序使用的是组合排序，快排，插入排序和堆排全用上了

	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left]+arr[right] == target {
			break
		}

		if arr[left]+arr[right] < target { // 小于target的情况，左边++ 变大
			left++
			continue
		}

		if arr[left]+arr[right] > target { // 相加大于target的情况，右边变小
			right--
			continue
		}
	}

	return []int{left, right}
}

// 使用Hashmap来实现时间复杂度 < O(N^2)
func twosumHash(arr []int, target int) []int {
	tmpMap := make(map[int]int)

	for i, val := range arr {
		if key, ok := tmpMap[target-val]; !ok {
			tmpMap[val] = i
		} else {
			return []int{i, tmpMap[key]}
		}
	}

	return []int{}
}

//func twosumHash(arr []int, target int) []int {
//	numMap := make(map[int][]int) // map[val][]index => 因为一个数组value相同的元素可能存在多个
//	for idx, val := range arr {
//		numMap[val] = append(numMap[val], idx)
//	}
//
//	for ele, ids := range numMap {
//		if val, ok := numMap[target-ele]; ok {
//			// 存在这个元素，看看是否相等，如果相等，必然有且只有两个
//
//			if len(ids) == 1 && len(val) == 1 && ids[0] == val[0] { // 两个元素是同一个
//				continue
//			}
//
//			if len(ids) == 1 && len(val) == 1 && ids[0] != val[0] { // 两个元素不是同一个且满足条件
//				return []int{ids[0], val[0]}
//			}
//
//			if len(ids) == 2 && len(val) == 2 {
//				return ids
//			}
//		}
//	}
//
//	return []int{}
//}

// 使用Hashmap来实现时间复杂度 < O(N^2)
func twosumHashV1(arr []int, target int) []int {
	numMap := make(map[int]int) // map[val][]index => 因为一个数组value相同的元素可能存在多个
	for idx, val := range arr {
		if index, ok := numMap[target-val]; ok {
			return []int{index, idx}
		}

		numMap[val] = idx
	}

	return []int{}
}

// 使用Hashmap来实现时间复杂度 < O(N^2)
func twosumBubble(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

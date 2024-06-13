package leetcode

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
	numMap := make(map[int][]int) // map[val][]index => 因为一个数组value相同的元素可能存在多个
	for idx, val := range arr {
		numMap[val] = append(numMap[val], idx)
	}

	for ele, ids := range numMap {
		if val, ok := numMap[target-ele]; ok {
			// 存在这个元素，看看是否相等，如果相等，必然有且只有两个

			if len(ids) == 1 && len(val) == 1 && ids[0] == val[0] { // 两个元素是同一个
				continue
			}

			if len(ids) == 1 && len(val) == 1 && ids[0] != val[0] { // 两个元素不是同一个且满足条件
				return []int{ids[0], val[0]}
			}

			if len(ids) == 2 && len(val) == 2 {
				return ids
			}
		}
	}

	return []int{}
}

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

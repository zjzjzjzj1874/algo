package easy

// 219. 存在重复元素 II
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
// 示例 2：
//
// 输入：nums = [1,0,1,1], k = 1
// 输出：true
// 示例 3：
//
// 输入：nums = [1,2,3,1,2,3], k = 2
// 输出：false
//
// 提示：
//
// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109
// 0 <= k <= 105

// 解题：hash表
func containsNearbyDuplicate(nums []int, k int) bool {
	// hash表和滑动窗口
	numMap := make(map[int]int)

	for i, num := range nums {
		idx, ok := numMap[num] // 是否存在相同元素
		if !ok {
			numMap[num] = i
			continue
		}

		if i-idx <= k {
			return true
		} else {
			numMap[num] = i // 更新为最近的一个
		}
	}

	return false
}

// 解题：滑动窗口
func containsNearbyDuplicateSlide(nums []int, k int) bool {
	n := len(nums)

	slow := 0
	for ; slow < n; slow++ {
		for fast := slow + 1; fast <= slow+k && fast < n; fast++ {
			if nums[slow] == nums[fast] { // 因为在k的滑窗查找的，fast-slow<=k一定满足
				return true
			}
		}
	}

	return false
}

// 解题：滑动窗口+hash表
func containsNearbyDuplicateHash(nums []int, k int) bool {
	cache := make(map[int]struct{})

	for i := range nums {
		if i > k {
			delete(cache, nums[i-k-1])
		}

		if _, ok := cache[nums[i]]; ok {
			return true
		}
		cache[nums[i]] = struct{}{}
	}

	return false
}

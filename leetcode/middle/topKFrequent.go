package middle

import "sort"

// 47. 前 K 个高频元素
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:
//
// 输入: nums = [1], k = 1
// 输出: [1]
//
// 提示：
//
// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

// 解题：
func topKFrequent(nums []int, k int) []int {
	// 前K个，最多有k个，还需要去重 hash表去重
	nMap := make(map[int]int) // map[num]count

	for i := range nums {
		nMap[nums[i]]++
	}

	ans := make([]int, 0, len(nMap))
	for key := range nMap {
		ans = append(ans, key)
	}

	sort.Slice(ans, func(i, j int) bool {
		return nMap[ans[i]] > nMap[ans[j]]
	})

	return ans[:k]
}

//  解题思路：使用hash表，先存储

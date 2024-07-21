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

// 解题思路：使用hash表，存储元素和出现次数，然后排序
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

// 解题：计数排序,可以找出现频率topK和minK的值
func topKFrequentWithCount(nums []int, k int) []int {
	// 前K个，最多有k个，还需要去重 hash表去重
	nMap := make(map[int]int) // map[num]count

	for _, num := range nums {
		nMap[num]++
	}
	// 计数排序，创建一个大小为n+1的数组，cnt[val]表示出现x次的数字的个数
	cnt := make([]int, len(nums)+1)
	for _, val := range nMap {
		cnt[val]++
	}

	i := len(nums)
	j := 0 // 记录累计频率计数，
	// 从高到低遍历 cnt 数组，找到累计频率达到 k 的频率值 i。
	for ; i >= 0; i-- {
		j = j + cnt[i]
		if j >= k { // 当j累计到k一样的时候，停止计数，并停止循环
			break
		}
	}

	ans := make([]int, 0, k)
	for k, v := range nMap {
		if v >= i {
			ans = append(ans, k)
		}
	}

	return ans
}

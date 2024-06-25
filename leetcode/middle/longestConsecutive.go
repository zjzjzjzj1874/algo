package middle

// 128. 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
// 提示：
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

// 解题：1.先把所有元素装入hashmap；2.遍历，每个元素++，判读是否存在，不存在break
func longestConsecutive(nums []int) int {
	numMap := make(map[int]struct{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	max := 0
	for num := range numMap {
		if _, ok := numMap[num-1]; ok { // 当上一个元素存在，不做循环，避免非第一个连续的也加入判断
			continue
		}
		temp := 1
		for {
			if _, ok := numMap[temp+num]; ok {
				temp++
			} else {
				break
			}
		}
		if temp > max {
			max = temp
		}
	}

	return max
}

package easy

import "sort"

// https://leetcode.cn/problems/3sum/description/
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4] => [-1,-1,0,1,2] => -1 + 2 > 0
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。
// 提示：
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105

// i + j + k = 0; i + k = -j => 理解为：两个数找等于第三个数相反数的数；i+k > -j ; => i+k+j > 0，j右移；i+k+j<0，j左移

func threesum(nums []int) [][]int {
	sort.Ints(nums) // 先排序

	// 找出第一个>0的Index
	firstZeroIdx := 0
	for idx, i := range nums {
		if i >= 0 {
			firstZeroIdx = idx
			break
		}
	}

	left := 0                  // 左指针
	right := len(nums) - 1     // 右指针
	res := make([][]int, 0, 4) // 先指定容量为4，超过自动扩容
	if firstZeroIdx == 0 && nums[1] == 0 && nums[2] == 0 {
		res = append(res, []int{0, 0, 0})
		if len(nums) == 3 {
			return res
		}
		left = 3
		for i := 3; i < len(nums); i++ {
			if nums[i] >= 0 {
				firstZeroIdx = i
				break
			}
		}
	}

	for left < firstZeroIdx && right > firstZeroIdx { // 停止条件，左指针或者右指针走过了第一个零指针
		for i := left + 1; i < right; i++ {
			if nums[left]+nums[right]+nums[i] == 0 { // 三个数加起来已经==0了，i尝试再往右移，不排除arr[i]和arr[i+1] == 0
				res = append(res, []int{nums[left], nums[right], nums[i]})
				if nums[left]+nums[right] == 0 { // 两数已经为0，终止
					break
				}
			}
		}
		if nums[left]+nums[right] >= 0 && left < firstZeroIdx { // 两个加起来大于0，游标必须在第一个0左边
			for right > firstZeroIdx {
				oldRight := right
				right--
				if nums[oldRight] != nums[right] {
					break
				}
			}
			continue
		}
		if nums[left]+nums[right] < 0 && right > firstZeroIdx { // 两个加起来大于0，游标必须在第一个0左边
			for left < firstZeroIdx {
				oldLeft := left
				left++
				if nums[oldLeft] != nums[left] {
					break
				}
			}
			left++
			continue
		}
		//	都不满足上述条件，left++
		left++
	}

	return res
}

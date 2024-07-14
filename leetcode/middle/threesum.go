package middle

import "sort"

// 15. 三数之和 https://leetcode.cn/problems/3sum/description/
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

// 解题：i + j + k = 0; i + k = -j => 理解为：两个数找等于第三个数相反数的数；i+k > -j ; => i+k+j > 0，j右移；i+k+j<0，j左移

//	func threeSum(nums []int) [][]int {
//		sort.Ints(nums) // 先对数组进行排序
//
//		result := make([][]int, 0, 4)
//		for first := 0; first < len(nums); first++ {
//			if first > 0 && nums[first] == nums[first-1] { // 第一个数和上一个数相同时，不做任何操作
//				continue
//			}
//
//			target := -nums[first] // 退化为双指针问题
//			third := len(nums) - 1
//			second := first + 1
//			for second < third {
//				if second > first+1 && nums[second] == nums[second-1] { // 第二个数和上一个数相同时，不做任何操作
//					second++
//					continue
//				}
//				if nums[second]+nums[third] > target {
//					third--
//					continue
//				} else if nums[second]+nums[third] < target {
//					second++
//					continue
//				}
//
//				if nums[second]+nums[third] == target {
//					result = append(result, []int{nums[first], nums[second], nums[third]})
//					second++
//				}
//			}
//		}
//
//		return result
//	}
func threeSum(nums []int) [][]int {
	// 双指针+两数之和
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums) // 对数组排序

	ans := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 和前一个数重复，直接跳过，避免重复计数
			continue
		}
		sum := -1 * nums[i]
		// 退化成两数之和

		x := i + 1
		y := n - 1

		for x < y {
			if x > i+1 && nums[x] == nums[x-1] {
				x++
				continue // 重复，避免重复计数
			}

			if nums[x]+nums[y] == sum {
				ans = append(ans, []int{nums[i], nums[x], nums[y]})
				x++
			} else if nums[x]+nums[y] > sum {
				y--
			} else {
				x++
			}
		}
	}

	return ans
}

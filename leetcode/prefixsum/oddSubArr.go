package prefixsum

// 1524. 和为奇数的子数组数目
// 给你一个整数数组 arr 。请你返回和为 奇数 的子数组数目。
//
// 由于答案可能会很大，请你将结果对 10^9 + 7 取余后返回。
//
// 示例 1：
//
// 输入：arr = [1,3,5]
// 输出：4
// 解释：所有的子数组为 [[1],[1,3],[1,3,5],[3],[3,5],[5]] 。
// 所有子数组的和为 [1,4,9,3,8,5].
// 奇数和包括 [1,9,3,5] ，所以答案为 4 。
// 示例 2 ：
//
// 输入：arr = [2,4,6]
// 输出：0
// 解释：所有子数组为 [[2],[2,4],[2,4,6],[4],[4,6],[6]] 。
// 所有子数组和为 [2,6,12,4,10,6] 。
// 所有子数组和都是偶数，所以答案为 0 。
// 示例 3：
//
// 输入：arr = [1,2,3,4,5,6,7]
// 输出：16
// 示例 4：
//
// 输入：arr = [100,100,99,99]
// 输出：4
// 示例 5：
//
// 输入：arr = [7]
// 输出：1
//
// 提示：
//
// 1 <= arr.length <= 10^5
// 1 <= arr[i] <= 100

// 靠解题过日子的菜鸟选手
func numOfSubArrays(arr []int) int {
	// 奇数和的前缀和
	sum := 0
	odd := 0
	even := 1
	subArr := 0

	for _, num := range arr {
		sum += num
		if sum%2 == 0 {
			// 偶数
			subArr += odd
			even++
		} else {
			// 奇数
			subArr += even
			odd++
		}
	}

	return subArr % 1000000007
}

// 暴力解题，超时
func numOfSubArraysWith(arr []int) int {
	// 奇数和的前缀和

	sum := 0
	nums := make([]int, 0, len(arr))

	for _, num := range arr {
		sum += num
		nums = append(nums, sum)
	}

	ans := 0 // 一共出现的次数
	for i, num := range nums {
		if num%2 != 0 {
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			if (num+nums[j])%2 != 0 {
				ans++
			}
		}
	}

	return ans
}

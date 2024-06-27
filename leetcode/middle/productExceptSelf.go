package middle

// 238. 除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//示例 1:
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//示例 2:
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
//提示：
//2 <= nums.length <= 105
//-30 <= nums[i] <= 30
//保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内
//进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

// 解题：申明两个数组，将左右数组的乘积缓存起来，注意如何区分这两个的边界
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right, ans := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}

// 解题：申明返回的数组作为上面的Left数组，然后Right数组动态构造
func productExceptSelfV1(nums []int) []int {
	n := len(nums)
	ans := make([]int, len(nums))

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	r := 1 // 默认right是左边的那个乘积
	for i := n - 1; i >= 0; i-- {
		ans[i] = r * ans[i]
		r *= nums[i]
	}

	return ans
}

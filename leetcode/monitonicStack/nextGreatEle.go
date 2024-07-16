package monitonicStack

// TODO 这个题没过，方法有问题，要使用排列或者反转，暂时不做，先把单调栈的题巩固一下
// 556. 下一个更大元素 III
// 类似题目：
// 31. 下一个排列
// 7. 整数反转
// 给你一个正整数 n ，请你找出符合条件的最小整数，其由重新排列 n 中存在的每位数字组成，并且其值大于 n 。如果不存在这样的正整数，则返回 -1 。
// 注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，但不是 32 位整数 ，同样返回 -1 。
// 示例 1：
// 输入：n = 12
// 输出：21
// 示例 2：
//
// 输入：n = 21
// 输出：-1
//
// 提示：
// 1 <= n <= 231 - 1

// 解题：先把n转换成数组
func nextGreaterElement556(n int) int {
	nums := make([]int, 0, 4)
	for i := n; i > 0; i /= 10 {
		j := i % 10
		nums = append(nums, j)
	}

	length := len(nums)
	ans := make([]int, length)
	stack := make([]int, 0, length)

	for i := length - 1; i >= 0; i-- {
		// stack不为空且nums[i]比栈顶的元素值大
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	idx := -1
	for i := 1; i < length; i++ {
		if ans[i] > 0 && ans[i] > nums[i] {
			idx = i
			break
		}
	}
	if idx == -1 {
		return idx
	}
	// 这里做切割
	res := 0
	carry := 1
	for i := idx; i >= 0; i-- {
		res = res + nums[i]*carry
		carry *= 10
	}
	for i := idx + 1; i < length; i++ {
		res = res + nums[i]*carry
		carry *= 10
	}

	return res
}

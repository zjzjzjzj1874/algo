package easy

// 9. 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数
// 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 例如，121 是回文，而 123 不是。
//
// 示例 1：
//
// 输入：x = 121
// 输出：true
// 示例 2：
//
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3：
//
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
//
// 提示：
//
// -231 <= x <= 231 - 1
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？

// 解题：转成数组
func isPalindromeNumWithNums(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	nums := make([]int, 0, 4)

	for d := x; d > 0; d /= 10 {
		nums = append(nums, d%10)
	}

	n := len(nums)
	for i := 0; i < n/2; i++ {
		if nums[i] != nums[n-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeWithNum(x int) bool {
	switch x {

	}
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 { // 负数和以0结尾的数必然不是回文数
		return false
	}

	reverse := 0

	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	// x == reverse: 只有四位数(偶数)，1221，上面跑完，x = 12, reverse = 12;
	// x == reverse/10: 只有五位数(奇数)，12321，上面跑完：x = 12, reverse = 123, 但是第三位的3不影响是否是回文数，直接/10删除掉即可
	return x == reverse || x == reverse/10
}

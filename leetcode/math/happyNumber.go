package math

// 202. 快乐数
// 编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
//
// 示例 1：
//
// 输入：n = 19
// 输出：true
// 解释：
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
// 示例 2：
//
// 输入：n = 2
// 输出：false
//
// 提示：
//
// 1 <= n <= 231 - 1

// 解题：暴力解法
func isHappy(n int) bool {
	valMap := make(map[int]bool)

	for n != 1 {
		if valMap[n] { // 这个数已经存在，表示循环了，返回false
			return false
		}
		valMap[n] = true

		t := n
		n = 0
		for t >= 10 {
			one := t % 10 // 个位数
			n += one * one
			t /= 10
		}
		n += t * t
	}

	return true
}

// 解题：快慢指针，快乐数：fast先到达1；不是快乐数：快慢指针在某个数相遇
func isHappyWithTwoPtr(n int) bool {
	calPowSum := func(n int) int {
		t := n
		sum := 0
		for t >= 10 {
			one := t % 10 // 个位数
			sum += one * one
			t /= 10
		}
		sum += t * t

		return sum
	}

	slow := n
	fast := calPowSum(n)
	for fast != 1 && fast != slow {

		slow = calPowSum(slow)
		fast = calPowSum(calPowSum(fast))
	}

	return fast == 1
}

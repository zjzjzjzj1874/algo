package middle

import "math"

// 633. 平方数之和
// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
//
// 示例 1：
//
// 输入：c = 5
// 输出：true
// 解释：1 * 1 + 2 * 2 = 5
// 示例 2：
//
// 输入：c = 3
// 输出：false
//
// 提示：
//
// 0 <= c <= 2^31 - 1

// 解题：双指针，这里可以使用golang的库函数了
func judgeSquareSum(c int) bool {
	if c < 3 {
		return true
	}
	l := 0
	r := int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}

	return false
}

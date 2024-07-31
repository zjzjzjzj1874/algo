package easy

import "math"

// LCR 001. 两数相除
// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
//
// 注意：
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1
//
// 示例 1：
//
// 输入：a = 15, b = 2
// 输出：7
// 解释：15/2 = truncate(7.5) = 7
// 示例 2：
//
// 输入：a = 7, b = -3
// 输出：-2
// 解释：7/-3 = truncate(-2.33333..) = -2
// 示例 3：
//
// 输入：a = 0, b = 1
// 输出：0
// 示例 4：
//
// 输入：a = 1, b = 1
// 输出：1
//
// 提示:
//
// -231 <= a, b <= 231 - 1
// b != 0
//
// 注意：本题与主站 29 题相同：https://leetcode-cn.com/problems/divide-two-integers/

func divide(a int, b int) int {
	if a == 0 {
		return 0
	}

	if b == 1 {
		return a
	} else if b == -1 {
		if a > math.MinInt32 {
			return -a
		}
		return math.MaxInt32
	}

	neg := false
	// 把a，b都转成负数
	if a > 0 {
		a = -a
		neg = !neg // 取反
	}
	if b > 0 {
		b = -b
		neg = !neg // 取反
	}
	meta := []int{b}
	for y := b; y >= a-y; {
		y += y
		meta = append(meta, y)
	}
	ans := 0
	for i := len(meta) - 1; i >= 0; i-- {
		if meta[i] >= a {
			ans |= 1 << i
			a -= meta[i]
		}
	}

	if neg {
		ans = -ans
	}

	return ans
}

// 解题：辗转相减法：但是会超时，必须进行优化
func divideTimeout(a int, b int) int {
	if a == 0 {
		return 0
	}

	if b == 1 {
		return a
	} else if b == -1 {
		if a > math.MinInt32 {
			return -a
		}
		return math.MaxInt32
	}

	times := 0 // 辗转相减的次数

	if a > 0 && b > 0 {
		for a > b {
			times++
			a -= b
		}
	} else if a < 0 && b < 0 {
		for a < b {
			times++
			a -= b
		}
	} else if a > 0 && b < 0 {
		b = -b
		for a > b {
			times++
			a -= b
		}
		times = -times
	} else { // a < 0 && b > 0
		a = -a
		for a > b {
			times++
			a -= b
		}
		times = -times
	}

	return times
}

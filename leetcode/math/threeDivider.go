package math

import "math"

// 1952. 三除数
// 给你一个整数 n 。如果 n 恰好有三个正除数 ，返回 true ；否则，返回 false 。
//
// 如果存在整数 k ，满足 n = k * m ，那么整数 m 就是 n 的一个 除数 。
//
// 示例 1：
//
// 输入：n = 2
// 输出：false
// 解释：2 只有两个除数：1 和 2 。
// 示例 2：
//
// 输入：n = 4
// 输出：true
// 解释：4 有三个除数：1、2 和 4 。
//
// 提示：
//
// 1 <= n <= 104
func isThree(n int) bool {
	if n < 4 {
		return false
	}

	num := 2

	for i := n / 2; i > 1; {
		if n%i == 0 {
			num++
		}
		if num > 3 {
			return false
		}
		pow := int(math.Sqrt(float64(n)))
		if i/2 < pow && i != pow {
			i = pow
		} else {
			i /= 2
		}
	}

	return num == 3
}

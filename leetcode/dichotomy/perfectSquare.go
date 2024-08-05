package dichotomy

// 367. 有效的完全平方数
// 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
// 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
//
// 不能使用任何内置的库函数，如  sqrt 。
//
// 示例 1：
//
// 输入：num = 16
// 输出：true
// 解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数。
// 示例 2：
//
// 输入：num = 14
// 输出：false
// 解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。
//
// 提示：
//
// 1 <= num <= 231 - 1

// 解题：二分法
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	l := 1
	r := num
	//for l < r {
	//	mid := l + (r-l+1)/2
	//	if float64(mid) == float64(num)/float64(mid) {
	//		return true
	//	} else if float64(mid) > float64(num)/float64(mid) {
	//		r = mid - 1
	//	} else {
	//		l = mid
	//	}
	//}

	for l < r {
		mid := l + (r-l+1)/2
		if num%mid == 0 && mid == num/mid { // 防止溢出同时防止丢失精度
			return true
		} else if mid > (num)/(mid) {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return false
}

// 暴力枚举
func isPerfectSquareWithIterate(num int) bool {
	if num == 1 {
		return true
	}
	for x := 1; x*x <= num; x++ {
		if x*x == num {
			return true
		}
	}

	return false
}

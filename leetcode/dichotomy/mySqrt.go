package dichotomy

// 69. x 的平方根
// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
// 示例 1：
//
// 输入：x = 4
// 输出：2
// 示例 2：
//
// 输入：x = 8
// 输出：2
// 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
//
// 提示：
//
// 0 <= x <= 231 - 1

// 解题：这个使用二分法，不断缩小范围
func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	if x <= 3 {
		return 1
	}

	l := 1
	r := x

	for l < r {
		mid := l + (r-l+1)/2 // +1表示：如果两个数相邻后，找右边的数作为中点，不然这个for跳不出去

		// 使用除法防止溢出
		if mid == x/mid {
			return mid
		}
		if mid > x/mid {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l
}
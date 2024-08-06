package math

// 计算两数的最小公倍数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// 计算最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

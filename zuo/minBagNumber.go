package zuo

// 最小带子数量
// 小虎去附近的商店买苹果，奸诈的商贩使用了捆绑交易，只提供6个每袋和8个每袋的包装包装不可拆分。可是小虎现在只想购买恰好N个苹果，小虎想购买尽量少的袋数方便携带。
// 如果不能购买恰好N个苹果，小虎将不会购买。输入一个整数N，表示小虎想购买的个苹果，返回最小使用多少袋子。如果无论如何都不能正好装下，返回-1。
// 解题：这个题主要用贪心：企图全部用8类型的带子，如果8类型不满足，这个时候减少8类型的带子数量，当减少的数量大于他俩的最小公倍数还不能搞定，直接返回-1
func minBagNumber(n int) int {
	mod := n % 8
	if mod == 0 {
		return n / 8
	}
	if mod == 6 { // 如果刚好余数==6，则刚好这么多
		return n/8 + 1
	}

	lc := lcm(8, 6) // 求最小公倍数
	i := 0          // 表示8的带子，每次分一点出去，尝试用6的带子解决
	for n/8 >= i && 8*i+mod < lc {
		if (8*i+n%8)%6 == 0 {
			return (8*i+n%8)/6 + n/8 - i
		}
		i++
	}

	return -1
}

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

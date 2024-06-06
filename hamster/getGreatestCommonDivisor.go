package hamster

// 求两数的最大公约数

func getGreatestCommonDivisor(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	if a == b {
		return a
	}

	big, small := 0, 0
	if a > b {
		big = a
		small = b
	} else {
		big = b
		small = a
	}

	remainder := small
	common := 0

	for remainder != 0 {
		temp := big % small
		if temp == 0 {
			return small
		}

		big = small
		small = temp
		remainder = temp
	}

	return common
}

// 使用递归，外部判断过a，b的情况了，这里默认不会出现为0的情况
func getGreatestCommonDivisorV2(a, b int) int {
	big, small := 0, 0
	if a > b {
		big = a
		small = b
	} else {
		big = b
		small = a
	}

	if big%small == 0 {
		return small
	}

	return getGreatestCommonDivisorV2(small, big%small)
}

// 使用递归，V2使用取余来整，但是两个数较大时，取模性能较低，使用更相减损术(九章算术)来实现；
func getGreatestCommonDivisorV3(a, b int) int {
	big, small := 0, 0
	if a > b {
		big = a
		small = b
	} else {
		big = b
		small = a
	}

	if big-small == 0 {
		return small
	}

	return getGreatestCommonDivisorV3(small, big-small)
}

// TODO 相除法和相减法相结合
//
//	当a和b均为偶数时，gcd(a,b) = 2×gcd(a/2, b/2) = 2×gcd(a>>1,b>>1)。
//	当a为偶数，b为奇数时，gcd(a,b) = gcd(a/2,b) = gcd(a>>1,b)。
//	当a为奇数，b为偶数时，gcd(a,b) = gcd(a,b/2) = gcd(a,b>>1)。
//	当a和b均为奇数时，先利用更相减损术运算一次，gcd(a,b) = gcd(b,a- b)，此时a-b必然是偶数，然后又可以继续进行移位运算。
func getGreatestCommonDivisorV4(a, b int) int {
	big, small := 0, 0
	if a > b {
		big = a
		small = b
	} else {
		big = b
		small = a
	}

	if big-small == 0 {
		return small
	}

	return getGreatestCommonDivisorV4(small, big-small)
}

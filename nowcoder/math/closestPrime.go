package main

import "fmt"

// HJ60 查找组成一个偶数最接近的两个素数
// 描述
// 任意一个偶数（大于2）都可以由2个素数组成，组成偶数的2个素数有很多种情况，本题目要求输出组成指定偶数的两个素数差值最小的素数对。
//
// 数据范围：输入的数据满足4≤n≤1000
// 输入描述：
// 输入一个大于2的偶数
//
// 输出描述：
// 从小到大输出两个素数
//
// 示例1
// 输入：
// 20
// 复制
// 输出：
// 7
// 13
// 复制
// 示例2
// 输入：
// 4
// 复制
// 输出：
// 2
// 2
func main() {
	n := 0
	_, _ = fmt.Scan(&n)

	for i := n / 2; i > 0; i-- {
		if isPrime(i) && isPrime(n-i) {
			fmt.Printf("%d\n%d\n", i, n-i)
			break
		}
	}
}
func isPrime(num int) (ret bool) {
	if num == 2 {
		return true
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

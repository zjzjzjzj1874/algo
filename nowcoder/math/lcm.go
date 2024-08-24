package main

import "fmt"

// HJ108 求最小公倍数
// 描述
// 正整数A和正整数B 的最小公倍数是指 能被A和B整除的最小的正整数值，设计一个算法，求输入A和B的最小公倍数。
//
// 数据范围：1≤a,b≤100000
// 输入描述：
// 输入两个正整数A和B。
//
// 输出描述：
// 输出A和B的最小公倍数。
//
// 示例1
// 输入：
// 5 7
// 复制
// 输出：
// 35
// 复制
// 示例2
// 输入：
// 2 4
// 复制
// 输出：
// 4
func lcm() {
	//func main() {
	a := 0
	b := 0
	_, _ = fmt.Scan(&a, &b)
	ab := gcd(a, b)

	fmt.Println(a * b / ab)
}

// 求最大公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

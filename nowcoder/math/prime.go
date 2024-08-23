package main

import "fmt"

// HJ6 质数因子
// 描述
// 功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）
// 数据范围：1≤n≤2×10^9+14
// 输入描述：
// 输入一个整数
// 输出描述：
// 按照从小到大的顺序输出它的所有质数的因子，以空格隔开。
// 示例1
// 输入：
// 180
// 复制
// 输出：
// 2 2 3 3 5
func prime() {
	//func main() {
	n := 0
	for {
		_, _ = fmt.Scan(&n)
		if n < 4 {
			fmt.Println(n)
			continue
		}

		i := 2
		for i*i <= n {
			if n%i == 0 {
				fmt.Printf("%d ", i)
				n /= i
			} else {
				i++
			}
		}
		fmt.Printf("%d \n", n)
	}
}

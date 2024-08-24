package main

import "fmt"

// HJ10 字符个数统计
// 描述
// 编写一个函数，计算字符串中含有的不同字符的个数。字符在 ASCII 码范围内( 0~127 ，包括 0 和 127 )，换行表示结束符，不算在字符里。不在范围内的不作统计。多个相同的字符只计算一次
// 例如，对于字符串 abaca 而言，有 a、b、c 三种不同的字符，因此输出 3 。
//
// 数据范围：1≤n≤500
// 输入描述：
// 输入一行没有空格的字符串。
//
// 输出描述：
// 输出 输入字符串 中范围在(0~127，包括0和127)字符的种数。
//
// 示例1
// 输入：
// abc
// 复制
// 输出：
// 3
// 复制
// 示例2
// 输入：
// aaa
// 复制
// 输出：
// 1
func distinctChar() {
	//func main() {
	str := ""
	fmt.Scan(&str)

	//sMap := make(map[byte]struct{})
	//for i := range str {
	//	sMap[str[i]] = struct{}{}
	//}
	//
	//fmt.Println(len(sMap))

	fmt.Println(countUniqueChars(str))
}

func countUniqueChars(s string) int {
	// 创建一个长度为128的布尔数组,初始化为false
	charCount := make([]bool, 128)

	// 遍历字符串,将出现的字符对应位置设为true
	for i := range s {
		if s[i] <= 127 {
			charCount[s[i]] = true
		}
	}

	// 统计数组中为true的个数,即为不同字符的个数
	count := 0
	for i := range charCount {
		if charCount[i] {
			count++
		}
	}

	return count
}

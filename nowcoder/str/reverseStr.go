package main

import (
	"bufio"
	"fmt"
	"os"
)

// HJ106 字符逆序
// 描述
// 将一个字符串str的内容颠倒过来，并输出。
//
// 数据范围：1≤len(str)≤10000
// 输入描述：
// 输入一个字符串，可以有空格
//
// 输出描述：
// 输出逆序的字符串
//
// 示例1
// 输入：
// I am a student
// 复制
// 输出：
// tneduts a ma I
// 复制
// 示例2
// 输入：
// nowcoder
// 复制
// 输出：
// redocwon
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	con := sc.Text()

	n := len(con)
	for i := range con {
		fmt.Print(string(con[n-i-1]))
	}
	fmt.Print("\n")
}

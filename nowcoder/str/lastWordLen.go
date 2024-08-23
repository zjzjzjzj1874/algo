package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// HJ1 字符串最后一个单词的长度
// 描述
// 计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）
// 输入描述：
// 输入一行，代表要计算的字符串，非空，长度小于5000。
//
// 输出描述：
// 输出一个整数，表示输入字符串最后一个单词的长度。
//
// 示例1
// 输入：
// hello nowcoder
// 复制
// 输出：
// 8
// 复制
// 说明：
// 最后一个单词为nowcoder，长度为8

// 解题：如果要测试，把这个方法改成main即可
func lastWordLen() {
	//func main() {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	strs := strings.Fields(msg)
	last := strs[len(strs)-1]
	fmt.Println(len(last))
}

package main

import (
	"fmt"
	"sort"
)

// HJ27 查找兄弟单词
// 描述
// 定义一个单词的“兄弟单词”为：交换该单词字母顺序（注：可以交换任意次），而不添加、删除、修改原有的字母就能生成的单词。
// 兄弟单词要求和原来的单词不同。例如： ab 和 ba 是兄弟单词。 ab 和 ab 则不是兄弟单词。
// 现在给定你 n 个单词，另外再给你一个单词 x ，让你寻找 x 的兄弟单词里，按字典序排列后的第 k 个单词是什么？
// 注意：字典中可能有重复单词。
// 数据范围：1≤n≤1000 ，输入的字符串长度满足1≤len(str)≤10，1≤k<n
// 输入描述：
// 输入只有一行。 先输入字典中单词的个数n，再输入n个单词作为字典单词。 然后输入一个单词x 最后后输入一个整数k
// 输出描述：
// 第一行输出查找到x的兄弟单词的个数m 第二行输出查找到的按照字典顺序排序后的第k个兄弟单词，没有符合第k个的话则不用输出。
// 示例1
// 输入：
// 3 abc bca cab abc 1
// 复制
// 输出：
// 2
// bca
// 复制
// 示例2
// 输入：
// 6 cab ad abcd cba abc bca abc 1
// 复制
// 输出：
// 3
// bca
// 复制
// 说明：
// abc的兄弟单词有cab cba bca，所以输出3
// 经字典序排列后，变为bca cab cba，所以第1个字典序兄弟单词为bca
func main() {
	n := 0
	fmt.Scan(&n)

	bros := make([]string, 0, n-2)
	word := ""
	k := 0
	for i := 0; i < n+2; i++ {
		if i < n {
			fmt.Scan(&word)
			bros = append(bros, word)
		} else if i == n {
			fmt.Scan(&word)
		} else if i == n+1 {
			fmt.Scan(&k)
		}
	}

	realBro := make([]string, 0, len(bros))
	for _, bro := range bros {
		if isBros(word, bro) {
			realBro = append(realBro, bro)
		}
	}

	sort.Strings(realBro)
	fmt.Println(len(realBro))
	if k <= len(realBro) {
		fmt.Println(realBro[k-1])
	}
}

func isBros(word, bro string) bool {
	if len(word) != len(bro) || word == bro {
		return false
	}
	w := [127]int{}
	b := [127]int{}
	for i := 0; i < len(bro); i++ {
		w[word[i]-'a']++
		b[bro[i]-'a']++
	}

	return w == b
}

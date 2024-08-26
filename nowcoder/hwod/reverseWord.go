package hwod

import (
	"strings"
)

// 单词倒叙
// 题目描述
// 输入单行英文句子，里面包含英文字母，空格以及,.? 三种标点符号，请将句子内每个单词进行倒序，并输出倒序后的语句
// 输入描述
// 输入字符串 S，S 的长度 1≤N≤100
// 输出描述 输出逆序后的字符串

// 解题：库函数
func reverse(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}

	return strings.Join(words, " ")
}

// 解题：双指针之快慢指针
func reverseWith2Ptr(s string) string {
	n := len(s)
	slow := n - 1
	ans := ""
	for fast := n - 1; fast >= 0; fast-- {
		if s[slow] == ' ' && s[fast] != ' ' {
			// 找到词头
			slow = fast
		} else if s[slow] != ' ' && s[fast] == ' ' {
			// 找到完整的单词，开始记录
			ans += s[fast+1:slow+1] + " "
			slow = fast
		}
	}

	if s[0] != ' ' {
		// 第一个是个单词
		ans += s[0 : slow+1]
	} else {
		ans = ans[:len(ans)-1]
	}

	return ans
}

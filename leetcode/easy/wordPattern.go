package easy

import "strings"

// 290. 单词规律
// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
//
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true
// 示例 2:
//
// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
// 示例 3:
//
// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false
//
// 提示:
//
// 1 <= pattern.length <= 300
// pattern 只包含小写英文字母
// 1 <= s.length <= 3000
// s 只包含小写英文字母和 ' '
// s 不包含 任何前导或尾随对空格
// s 中每个单词都被 单个空格 分隔

// 解题：双射关系：使用Hash表存储
func wordPattern(pattern string, s string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		ch := pattern[i]
		if p2s[ch] != "" && p2s[ch] != word || (s2p[word] > 0 && s2p[word] != ch) {
			return false
		}

		p2s[ch] = word
		s2p[word] = ch
	}

	return true
}

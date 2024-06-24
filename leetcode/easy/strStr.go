package easy

// Leetcode-28：找出字符串中第一个匹配项的下标
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
// 示例 1：
// 输入：haystack = "sadbutsad", needle = "sad"
// 输出：0
// 解释："sad" 在下标 0 和 6 处匹配。
// 第一个匹配项的下标是 0 ，所以返回 0 。
// 示例 2：
// 输入：haystack = "leetcode", needle = "leeto"
// 输出：-1
// 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
// 提示：
// 1 <= haystack.length, needle.length <= 104
// haystack 和 needle 仅由小写英文字符组成

// 解题：先首字母相比较，当且仅当首字母相同时，继续走needle字母的循环
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack) && i+len(needle) <= len(haystack); i++ {
		if haystack[i] != needle[0] { // hayStack[i]和needle的首字节不相同，continue
			continue
		}
		allEqual := true
		for j := 1; j < len(needle); j++ {
			if haystack[i+j] != needle[j] { // needle的剩下字节和haystack剩下的比较
				allEqual = false
				break
			}
		}
		if allEqual {
			return i
		}
	}

	return -1
}

// 其他：KMP算法和Boyer-Moore算法
// kmp算法
// KMP算法的解释：http://jakeboxer.com/blog/2009/12/13/the-knuth-morris-pratt-algorithm-in-my-own-words/
// https://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
func strStrKMP(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

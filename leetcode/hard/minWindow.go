package hard

import "math"

// 76. 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
//
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s 和 t 由英文字母组成
//
// 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？

func minWindow(s string, t string) string {
	// 滑动窗口
	if s == t {
		return s
	}
	if len(s) < len(t) {
		return ""
	}

	tt := [58]int{} // t的数组
	ts := [58]int{}
	for _, b := range t {
		tt[b-'A']++
	}

	left := 0
	min := math.MaxInt
	ans := ""
	for right, b := range s {
		ts[b-'A']++

		for contains(ts, tt) { // 如果包含
			if min > right-left+1 {
				min = right - left + 1
				ans = s[left : right+1]
			}

			ts[s[left]-'A']--
			left++
		}
	}

	return ans
}

// a1是否包含a2
func contains(a1, a2 [58]int) bool {
	for i := 0; i < 58; i++ {
		if a1[i] < a2[i] {
			return false
		}
	}

	return true
}

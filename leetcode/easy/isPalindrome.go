package easy

import "strings"

// 125. 验证回文串
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
// 示例 1：
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
// 示例 2：
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。
// 示例 3：
// 输入：s = " "
// 输出：true
// 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
// 由于空字符串正着反着读都一样，所以是回文串。
//
// 提示：
//
// 1 <= s.length <= 2 * 105
// s 仅由可打印的 ASCII 字符组成

// 解题：先将s转成小写字母，去除特殊符号和数字，然后使用双指针
//
//	func isPalindrome(s string) bool {
//		if s == "" {
//			return true
//		}
//
//		ns := strings.Builder{}
//		for i := range s {
//			ns.Grow(1)
//			c := s[i]
//			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
//				ns.WriteByte(s[i])
//			}
//		}
//		nstr := strings.ToLower(ns.String())
//		n := len(nstr)
//		left := 0
//		right := n - 1
//		for left < right {
//			if nstr[left] != nstr[right] {
//				return false
//			}
//			left++
//			right--
//		}
//
//		return true
//	}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	str := ""
	for i := range s {
		if isAscii(s[i]) {
			str += string(s[i])
		}
	}

	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeWithO1(s string) bool {
	if len(s) <= 1 {
		return true
	}

	s = strings.ToLower(s)
	n := len(s)
	left := 0
	right := n - 1

	for left < right {
		if left < right && !isAscii(s[left]) {
			left++
			continue
		}
		if left < right && !isAscii(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isAscii(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

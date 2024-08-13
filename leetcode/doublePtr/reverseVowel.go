package doublePtr

// 345. 反转字符串中的元音字母
// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//
// 示例 1：
//
// 输入：s = "hello"
// 输出："holle"
// 示例 2：
//
// 输入：s = "leetcode"
// 输出："leotcede"
//
// 提示：
//
// 1 <= s.length <= 3 * 105
// s 由 可打印的 ASCII 字符组成
func reverseVowels(s string) string {
	n := len(s)

	left := 0
	right := n - 1

	ans := make([]byte, n)
	for left <= right {
		// 两个都是元音，交换位置
		if isVowel(s[left]) && isVowel(s[right]) {
			ans[left] = s[right]
			ans[right] = s[left]
			left++
			right--
		} else if !isVowel(s[right]) {
			ans[right] = s[right]
			right--
		} else if !isVowel(s[left]) {
			ans[left] = s[left]
			left++
		}
	}

	return string(ans)
}

func isVowel(aim byte) bool {
	return aim == 'a' || aim == 'e' || aim == 'i' || aim == 'o' || aim == 'u' ||
		aim == 'A' || aim == 'E' || aim == 'I' || aim == 'O' || aim == 'U'
}

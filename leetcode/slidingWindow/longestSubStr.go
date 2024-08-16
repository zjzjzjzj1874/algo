package slidingWindow

// 2414. 最长的字母序连续子字符串的长度
// 字母序连续字符串 是由字母表中连续字母组成的字符串。换句话说，字符串 "abcdefghijklmnopqrstuvwxyz" 的任意子字符串都是 字母序连续字符串 。
//
// 例如，"abc" 是一个字母序连续字符串，而 "acb" 和 "za" 不是。
// 给你一个仅由小写英文字母组成的字符串 s ，返回其 最长 的 字母序连续子字符串 的长度。
//
// 示例 1：
//
// 输入：s = "abacaba"
// 输出：2
// 解释：共有 4 个不同的字母序连续子字符串 "a"、"b"、"c" 和 "ab" 。
// "ab" 是最长的字母序连续子字符串。
// 示例 2：
//
// 输入：s = "abcde"
// 输出：5
// 解释："abcde" 是最长的字母序连续子字符串。
//
// 提示：
//
// 1 <= s.length <= 105
// s 由小写英文字母组成
func longestContinuousSubstring(s string) (ans int) {
	// 解题：同样用滑动窗口来解决，l和r表示指针，r-l = s[r]-s[l],更新max(ans,r-l+1)的值; 不相等时，移动l到r的位置
	l := 0

	for r := 0; r < len(s); r++ {
		if s[r]-s[l] == byte(r-l) {
			ans = max(ans, r-l+1)
		} else {
			l = r
		}
	}
	return
}

package slidingWindow

// 424. 替换后的最长重复字符
// 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
//
// 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。
//
// 示例 1：
//
// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。
// 示例 2：
//
// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：
// 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
// 子串 "BBBB" 有最长重复字母, 答案为 4。
// 可能存在其他的方法来得到同样的结果。
//
// 提示：
//
// 1 <= s.length <= 105
// s 仅由大写英文字母组成
// 0 <= k <= s.length
func characterReplacement(s string, k int) int {

	n := len(s)

	alpha := [26]int{}
	res := 0 // 记录出现次数最多的元素
	l := 0
	for r := 0; r < n; r++ {
		alpha[s[r]-'A']++ // 记录出现频次

		res = max(res, alpha[s[r]-'A'])
		if r-l+1-res > k {
			alpha[s[l]-'A']--
			l++
		}
	}

	return n - l
}

package middle

// 438. 找到字符串中所有字母异位词
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
// 示例 1:
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//  示例 2:
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
// 提示:
// 1 <= s.length, p.length <= 3 * 104
// s 和 p 仅包含小写字母

// 解题：滑动窗口，小写字母，则对ASCII取模也可以；
func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)

	if sLen < pLen {
		return []int{}
	}

	pSub := [26]byte{} // p的子串
	sSub := [26]byte{} // s的子串
	res := make([]int, 0, 4)

	for i := range p {
		pSub[p[i]%26]++
		sSub[s[i]%26]++
	}

	if pSub == sSub {
		res = append(res, 0)
	}

	for i := 1; i <= sLen-pLen; i++ { // 注意边界，i从0还是从1开始，其边界都需要注意
		sSub[s[i-1]%26]--
		sSub[s[i+pLen-1]%26]++

		if sSub == pSub {
			res = append(res, i)
		}
	}

	return res
}

package slidingWindow

// 1456. 定长子串中元音的最大数目
// 给你字符串 s 和整数 k 。
//
// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
//
// 英文中的 元音字母 为（a, e, i, o, u）。
//
// 示例 1：
//
// 输入：s = "abciiidef", k = 3
// 输出：3
// 解释：子字符串 "iii" 包含 3 个元音字母。
// 示例 2：
//
// 输入：s = "aeiou", k = 2
// 输出：2
// 解释：任意长度为 2 的子字符串都包含 2 个元音字母。
// 示例 3：
//
// 输入：s = "leetcode", k = 3
// 输出：2
// 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
// 示例 4：
//
// 输入：s = "rhythms", k = 4
// 输出：0
// 解释：字符串 s 中不含任何元音字母。
// 示例 5：
//
// 输入：s = "tryhard", k = 4
// 输出：1
//
// 提示：
//
// 1 <= s.length <= 10^5
// s 由小写英文字母组成
// 1 <= k <= s.length

// 解题：滑动窗口，用一个字符串维护
// 定长滑窗套路
// 我总结成三步：入-更新-出。
//
// 入：下标为 i 的元素进入窗口，更新相关统计量。如果 i<k−1 则重复第一步。
// 更新：更新答案。一般是更新最大值/最小值。
// 出：下标为 i−k+1 的元素离开窗口，更新相关统计量。
func maxVowels(s string, k int) int {
	l := 0
	r := 0
	//a b c i i i d e f ;k == 3
	ans := 0
	tmp := 0
	for ; r < len(s); r++ {
		if isVowel(s[r]) {
			tmp++
		}
		if tmp > ans {
			ans = tmp
		}
		if r-l+1 == k {
			if isVowel(s[l]) {
				tmp--
			}
			l++
		}
	}

	return ans
}

func isVowel(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}

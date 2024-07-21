package middle

// LCR 014. 字符串的排列
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
//
// 示例 1：
//
// 输入: s1 = "ab" s2 = "eidbaooo"
// 输出: True
// 解释: s2 包含 s1 的排列之一 ("ba").
// 示例 2：
//
// 输入: s1= "ab" s2 = "eidboaoo"
// 输出: False
//
// 提示：
//
// 1 <= s1.length, s2.length <= 104
// s1 和 s2 仅包含小写字母
//
// 注意：本题与主站 567 题相同： https://leetcode-cn.com/problems/permutation-in-string/

// 解题：滑动窗口
func checkInclusion(s1 string, s2 string) bool {
	// 滑动窗口
	m1 := [26]int{}
	for i := range s1 {
		m1[s1[i]-'a']++
	}

	m2 := [26]int{}
	n1 := len(s1)

	for l, r := 0, 0; r < len(s2); r++ {
		m2[s2[r]-'a']++

		if r-l+1 == n1 {
			if m1 == m2 {
				return true
			}
			m2[s2[l]-'a']--
			l++
		}
	}

	return false
}

// 解题：滑动窗口
func checkInclusionWithLT(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}

	// 滑动窗口
	m1, m2 := [26]int{}, [26]int{}

	for i := 0; i < n1; i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	if m1 == m2 {
		return true
	}

	for i := n1; i < n2; i++ {
		m2[s2[i]-'a']++
		m2[s2[i-n1]-'a']--
		if m1 == m2 {
			return true
		}
	}

	return false
}

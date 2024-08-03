package middle

// 3. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 提示：
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成

// 解题：滑动窗口
func lengthOfLongestSubstring(s string) int {
	dup := make(map[byte]int) // hash表中字符出现的次数
	sub := 0                  // 不重复hash子串的截止位置
	longest := 0              // 最长不重复子串长度
	l := len(s)               // 字符串长度

	for idx := range s {
		if idx != 0 {
			delete(dup, s[idx-1]) // 除了第一次走，每次走外圈都删除最前面的一个
		}

		for sub < l {
			_, ok := dup[s[sub]]
			if !ok { // 不存在，加入hash，sub++
				dup[s[sub]]++
				sub++
			} else { // 存在，断开内循环，去外循环删除最前面的，一次删除，直到不存在时
				break
			}
		}

		longest = max(longest, len(dup))
	}

	return longest
}

//func lengthOfLongestSubstring(s string) int {
//	// 哈希集合，记录每个字符是否出现过
//	m := map[byte]int{}
//	n := len(s)
//	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
//	rk, ans := -1, 0
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			// 左指针向右移动一格，移除一个字符
//			delete(m, s[i-1])
//		}
//		for rk+1 < n && m[s[rk+1]] == 0 {
//			// 不断地移动右指针
//			m[s[rk+1]]++
//			rk++
//		}
//		// 第 i 到 rk 个字符是一个极长的无重复字符子串
//		ans = max(ans, rk-i+1)
//	}
//	return ans
//}

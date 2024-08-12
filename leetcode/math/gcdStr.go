package math

// 1071. 字符串的最大公因子 TODO unfinished
// 对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
// 给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
//
// 示例 1：
//
// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
// 示例 2：
//
// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"
// 示例 3：
//
// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""
//
// 提示：
//
// 1 <= str1.length, str2.length <= 1000
// str1 和 str2 由大写英文字母组成
func gcdOfStrings(s string, t string) string {
	ns := len(s)
	//nt := len(t)
	//
	//if nt > ns {
	//	return ""
	//}

	subt := minSub(t)
	if ns%len(subt) != 0 {
		return ""
	}

	for i := 0; i < ns; i += len(subt) {
		for j := 0; j < len(subt); j++ {
			if subt[j] != s[i+j] {
				return ""
			}
		}
	}

	return subt
}

func minSub(str string) string {
	n := len(str)

	if n&1 == 1 { // 奇数   0偶数
		return str
	}

	for i := n; i > 0; {
		tmp := str[:i]
		if i&1 == 1 { // 奇数，无法再分
			return tmp
		}

		// 对半砍
		i = i / 2
		for j := 0; j < i; j++ {
			if str[j] != str[i+j] { // 不满足，说明无法拼成功，返回tmp
				return tmp
			}
		}
	}

	return ""
}

package doublePtr

// 917. 仅仅反转字母
// 给你一个字符串 s ，根据下述规则反转字符串：
//
// 所有非英文字母保留在原有位置。
// 所有英文字母（小写或大写）位置反转。
// 返回反转后的 s 。
//
// 示例 1：
//
// 输入：s = "ab-cd"
// 输出："dc-ba"
// 示例 2：
//
// 输入：s = "a-bC-dEf-ghIj"
// 输出："j-Ih-gfE-dCba"
// 示例 3：
//
// 输入：s = "Test1ng-Leet=code-Q!"
// 输出："Qedo1ct-eeLg=ntse-T!"
//
// 提示
//
// 1 <= s.length <= 100
// s 仅由 ASCII 值在范围 [33, 122] 的字符组成
// s 不含 '\"' 或 '\\'
func reverseOnlyLetters(s string) string {
	ans := make([]byte, len(s))
	l := 0
	r := len(s) - 1
	for l <= r {
		if !isAlphabet(s[l]) {
			ans[l] = s[l]
			l++
			continue
		}
		if !isAlphabet(s[r]) {
			ans[r] = s[r]
			r--
			continue
		}
		ans[l] = s[r]
		ans[r] = s[l]
		l++
		r--
	}

	return string(ans)
}

func isAlphabet(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}

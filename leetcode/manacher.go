package leetcode

// Manacher算法是一种用于查找字符串中最长回文子串的线性时间复杂度算法。
// Preprocess函数:
//
// 在每个字符之间和两端添加#，将原始字符串转换为新字符串。例如，"babad"变为"#b#a#b#a#d#"。
// 这种转换使得处理奇数和偶数长度的回文子串变得一致。
// Manacher函数:
//
// p数组用于记录每个字符作为中心的回文半径。
// c是当前回文子串的中心，r是当前回文子串的右边界。
// 对于每个字符i，首先计算它的镜像位置mirror。
// 如果i在当前回文子串内，初始化p[i]为min(r-i, p[mirror])。
// 尝试从i开始扩展回文子串，更新p[i]。
// 如果新扩展的回文子串超过当前右边界r，更新c和r。
// 记录最长的回文子串的长度和中心位置。
//
// 测试函数
// 输入字符串"babad"，调用manacher函数，并输出最长回文子串。
// 时间复杂度: O(n)，其中n是预处理后字符串的长度，即原始字符串长度的两倍加一。Manacher算法只需线性时间即可完成。
// 空间复杂度: O(n)，用于存储预处理后的字符串和p数组

// 转换原始字符串，在每个字符之间以及开头和结尾添加“#”
// 这种转换有助于统一处理偶数和奇数回文
func preprocess(s string) string {
	t := "#"
	for _, ch := range s {
		t += string(ch) + "#"
	}
	return t
}

// Manacher算法寻找最长回文子串
func manacher(s string) string {
	t := preprocess(s) // 预处理字符串
	n := len(t)
	p := make([]int, n) // 用于记录每个字符作为中心的回文半径
	c := 0              // 当前回文子串的中心
	r := 0              // 当前回文子串的右边界
	maxLen := 0         // 扩出来的最大长度
	centerIndex := 0

	for i := 0; i < n; i++ {
		mirror := 2*c - i // i的镜像位置mirror

		if i < r {
			p[i] = min(r-i, p[mirror])
		}

		// 围绕当前中心展开
		for i+p[i]+1 < n && i-p[i]-1 >= 0 && t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}

		// 如果展开的回文较大，则更新中心和右边界
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}

		// 记录回文的最大长度
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	// 提取最长回文子字符串
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

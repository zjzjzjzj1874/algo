package hwod

//查找重复代码
//题目描述
//小明负责维护项目下的代码，需要查找出重复代码，用以支撑后续的代码优化，请你帮助 小明找出重复的代码，。
//重复代码查找方法:以字符串形式给定两行代码(字符串长度 1<length<= 100，由英文字母、数字和空格组成)，找出两行代码中的最长公共子串。
//注: 如果不存在公共子串，返回空字符串 输入描述
//输入的参数 text1, text2 分别表示两行代码 输出描述
//输出任一最长公共子串

func longestCommonSubString(s, t string) string {
	maxLen := 0
	maxEnd := 0
	m := len(s)
	n := len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					maxEnd = i
				}
			}
		}
	}

	if maxLen == 0 {
		return ""
	}
	
	return s[maxEnd-maxLen : maxEnd]
}

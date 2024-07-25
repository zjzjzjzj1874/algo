package backtrack

// LCR 086. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
// 回文串
// 。返回 s 所有可能的分割方案。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：
//
// 输入：s = "a"
// 输出：[["a"]]
//
// 提示：
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//注意：本题与主站 131 题相同： https://leetcode-cn.com/problems/palindrome-partitioning/

func partition(s string) (ans [][]string) {
	n := len(s)
	choice := make([]string, 0, n)

	// start表示开始位置,i表示当前位置
	var dfs func(i, start int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string{}, choice...))
			return
		}

		// 不选择当前，
		if i < n-1 {
			dfs(i+1, start)
		}

		if isPalindrome(s, start, i) {
			choice = append(choice, s[start:i+1])
			dfs(i+1, i+1) // 下一个子串从i+1开始
			choice = choice[:len(choice)-1]
		}
	}

	dfs(0, 0)
	return
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func partitionWithAns(s string) (ans [][]string) {
	n := len(s)
	choice := make([]string, 0, n)

	// i表示当前位置
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, choice...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				choice = append(choice, s[i:j+1])
				dfs(j + 1) // 下一个子串从j+1开始
				choice = choice[:len(choice)-1]
			}
		}
	}

	dfs(0)

	return
}

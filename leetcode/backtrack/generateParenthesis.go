package backtrack

import "fmt"

// LCR 085. 括号生成
// 正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
//
// 提示：
//
// 1 <= n <= 8
//
// 注意：本题与主站 22 题相同： https://leetcode-cn.com/problems/generate-parentheses/

// 解题：回溯
func generateParenthesis(n int) (ans []string) {
	choice := "" // 加入的选项
	left := n    // 左括号数量
	right := n   // 右括号数量

	var dfs func()
	dfs = func() {
		// 终止条件，左右括号都不剩余，直接加入结果集
		if left == 0 && right == 0 {
			ans = append(ans, fmt.Sprintf("%s", choice))
			return
		}

		// 加括号及剪枝
		if left > 0 {
			choice += "("
			left--
			dfs()
			choice = choice[:len(choice)-1] // 撤销左括号
			left++
		}
		if right > left { // 右括号的数量比左括号多，才能添加右括号
			choice += ")"
			right--
			dfs()
			choice = choice[:len(choice)-1]
			right++
		}
	}

	dfs()

	return ans
}

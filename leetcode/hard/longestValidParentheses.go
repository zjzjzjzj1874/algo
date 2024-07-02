package hard

// 32. 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号
// 子串
// 的长度。
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
// 示例 2：
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
// 示例 3：
//
// 输入：s = ""
// 输出：0
//
// 提示：
//
// 0 <= s.length <= 3 * 104
// s[i] 为 '(' 或 ')'

// 解题：贪心
func longestValidParentheses(s string) int {
	l, r := 0, 0
	n, max := len(s), 0

	// 从左向右计算
	for i := 0; i < n; i++ {
		ch := s[i]
		switch ch {
		case '(':
			l++
		case ')':
			r++
		}
		if l == r {
			if max < l+r {
				max = l + r
			}
		} else if r > l { // 如果右括号大于左括号了，说明失效了，重新开始
			r = 0
			l = 0
		}
	}

	// 从右向左
	l = 0
	r = 0
	for i := n - 1; i >= 0; i-- {
		ch := s[i]
		switch ch {
		case '(':
			l++
		case ')':
			r++
		}
		if l == r {
			if max < l+r {
				max = l + r
			}
		} else if l > r { // 如果左括号大于右括号了，说明失效了，重新开始
			r = 0
			l = 0
		}
	}

	return max
}

// 使用栈
func longestValidParenthesesStack(s string) int {
	maxAns := 0
	var stack []int           // 这里装的是左括号的坐标
	stack = append(stack, -1) // 这里初始化一个-1，表示上一个未匹配的右括号位置
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// i-stack[len(stack)-1] => 是坐标-坐标，表示的有效括号的长度
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

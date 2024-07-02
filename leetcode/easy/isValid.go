package easy

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
//
// 提示：
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
func isValid(s string) bool {
	if len(s)%2 == 1 { // 单数符号直接返回false
		return false
	}
	vMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)

	for i := range s {
		ch := s[i]

		val, ok := vMap[ch]
		if !ok { // 遍历的仍然是左括号，继续压入stack中
			stack = append(stack, ch)
			continue
		}

		// 已经有对应的右括号出现，判断val和stack中最后的一个值是否相等
		if len(stack) == 0 || stack[len(stack)-1] != val {
			return false
		}

		// 相等，则把stack的最后一个元素移除
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

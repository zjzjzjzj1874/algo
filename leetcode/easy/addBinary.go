package easy

import "fmt"

// 67. 二进制求和
// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
//
// 示例 1：
//
// 输入:a = "11", b = "1"
// 输出："100"
// 示例 2：
//
// 输入：a = "1010", b = "1011"
// 输出："10101"
//
// 提示：
//
// 1 <= a.length, b.length <= 104
// a 和 b 仅由字符 '0' 或 '1' 组成
// 字符串如果不是 "0" ，就不含前导零

func addBinary(a string, b string) string {
	na := len(a)
	nb := len(b)
	n := na
	if nb > n {
		n = nb
	}
	// n表示二进制的最长长度

	carry := 0 // 是否需要进位
	res := ""

	for i := 0; i < n; i++ {
		if i < na {
			carry += int(a[na-i-1] - '0') // 注意：这里应该是倒着做加法
		}
		if i < nb {
			carry += int(b[nb-i-1] - '0')
		}

		res = fmt.Sprintf("%d%s", carry%2, res)
		carry = carry / 2
	}

	if carry > 0 {
		res = "1" + res
	}

	return res
}

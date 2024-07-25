package backtrack

import (
	"fmt"
	"strconv"
)

// LCR 087. 复原 IP 地址
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
//
// 输入：s = "1111"
// 输出：["1.1.1.1"]
// 示例 4：
//
// 输入：s = "010010"
// 输出：["0.10.0.10","0.100.1.0"]
// 示例 5：
//
// 输入：s = "10203040"
// 输出：["10.20.30.40","102.0.30.40","10.203.0.40"]
//
// 提示：
//
// 0 <= s.length <= 3000
// s 仅由数字组成
//
// 注意：本题与主站 93 题相同：https://leetcode-cn.com/problems/restore-ip-addresses/

// 解题：回溯
func restoreIpAddresses(s string) (ans []string) {
	choice := make([]string, 0, 4)
	n := len(s)

	var dfs func(start int)

	dfs = func(start int) {
		if len(choice) == 4 { // 终止条件
			if start == n { // 所有的必须要用完，
				// 后置剪枝，即便长度为4，但是start没有走到最后，说明还有字符没用完，不能加入结果集
				// 拼接结果
				ans = append(ans, fmt.Sprintf("%s.%s.%s.%s", choice[0], choice[1], choice[2], choice[3]))
			}
			return
		}

		for i := start; i < n; i++ {
			if i != start && s[start] == '0' { // 前缀0只能单独存在，剪枝
				break
			}

			tmp := s[start : i+1]
			num, _ := strconv.Atoi(tmp)
			if num >= 0 && num <= 255 {
				choice = append(choice, tmp)

				dfs(i + 1)

				choice = choice[:len(choice)-1]
			} else {
				break
			}
		}
	}
	dfs(0)

	return
}

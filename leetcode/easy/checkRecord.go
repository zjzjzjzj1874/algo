package easy

import "strings"

// 51. 学生出勤记录 I
// 给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
// 'A'：Absent，缺勤
// 'L'：Late，迟到
// 'P'：Present，到场
// 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
// 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
// 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
// 如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：s = "PPALLP"
// 输出：true
// 解释：学生缺勤次数少于 2 次，且不存在 3 天或以上的连续迟到记录。
// 示例 2：
//
// 输入：s = "PPALLL"
// 输出：false
// 解释：学生最后三天连续迟到，所以不满足出勤奖励的条件。
//
// 提示：
//
// 1 <= s.length <= 1000
// s[i] 为 'A'、'L' 或 'P'
func checkRecord(s string) bool {
	a := 2
	l := 0

	for j := 0; j < len(s); j++ {
		if s[j] == 'A' {
			a--
		}
		if a <= 0 {
			return false
		}

		if s[j] == 'L' {
			l++
			if l >= 3 {
				return false
			}
		} else if s[j] != 'L' {
			l = 0 // 重置迟到天数
		}
	}

	return true
}

func checkRecordWithAI(s string) bool {
	return strings.Count(s, "A") < 2 && strings.Contains(s, "LLL")
}

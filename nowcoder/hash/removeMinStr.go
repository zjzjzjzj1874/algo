package main

import "fmt"

// HJ23 删除字符串中出现次数最少的字符
// 描述
// 实现删除字符串中出现次数最少的字符，若出现次数最少的字符有多个，则把出现次数最少的字符都删除。输出删除这些单词后的字符串，字符串中其它字符保持原来的顺序。
//
// 数据范围：输入的字符串长度满足 1≤n≤20，保证输入的字符串中仅出现小写字母
// 输入描述：
// 字符串只包含小写英文字母, 不考虑非法输入，输入的字符串长度小于等于20个字节。
//
// 输出描述：
// 删除字符串中出现次数最少的字符后的字符串。
//
// 示例1
// 输入：
// aabcddd
// 复制
// 输出：
// aaddd
func main() {
	s := ""
	_, _ = fmt.Scan(&s)

	countMap := make(map[byte]int)
	mini := len(s)
	for i := range s {
		countMap[s[i]]++
	}

	for _, val := range countMap {
		if val != 0 && val < mini {
			mini = val
		}
	}

	ans := make([]byte, 0, len(countMap))
	for i := range s {
		if countMap[s[i]] == mini {
			continue
		}
		ans = append(ans, s[i])
	}

	fmt.Println(string(ans))
}

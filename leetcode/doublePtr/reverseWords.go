package doublePtr

import (
	"fmt"
	"strings"
)

// 151. 反转字符串中的单词
// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
//
// 示例 1：
//
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 示例 2：
//
// 输入：s = "  hello world  "
// 输出："world hello"
// 解释：反转后的字符串中不能存在前导空格和尾随空格。
// 示例 3：
//
// 输入：s = "a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
//
// 提示：
//
// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词
//
// 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。

// 解题：使用原地算法，1.处理前后空格；2.反转所有的字母；3.根据单词来反转回去，得到答案
// 使用strings.Split失去了这道题的意义
func reverseWordsWithO1(s string) string {
	slow := len(s) - 1
	ans := ""

	for fast := len(s) - 1; fast >= 0; fast-- {
		if s[slow] == ' ' && s[fast] != ' ' { // 慢指针指向空格且快指针指向字母
			slow = fast // 则表明找到了单词的词头
		} else if s[fast] == ' ' && s[slow] != ' ' { // 快指针指向了空格，慢指针指向字母，说明找到了一个单词
			ans += s[fast+1:slow+1] + " " // 存入单词, 且加入空格
			slow = fast                   // 更新慢指针，指向空格，继续去找下一个词头
		}
	}

	if s[0] != ' ' { // s字符串是个单字母的单词
		ans += s[0 : slow+1]
	} else { // 表明没有单词了，把最后的空格移除
		ans = ans[:len(ans)-1]
	}

	return ans
}

func reverseString(s *[]byte, begin, end int) {
	for begin <= end {
		(*s)[begin], (*s)[end] = (*s)[end], (*s)[begin]
		begin++
		end--
	}
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	l := 0
	r := len(words) - 1

	for l <= r {
		if words[l] == "" {
			l++
			continue
		}
		if words[r] == "" {
			r--
			continue
		}
		tmp := words[l]
		words[l] = words[r]
		words[r] = tmp
		l++
		r--
	}

	ans := ""
	for i, word := range words {
		if word == "" {
			continue
		}
		if i == 0 || ans == "" {
			ans = word
		} else {
			ans = fmt.Sprintf("%s %s", ans, word)
		}
	}

	return ans
}

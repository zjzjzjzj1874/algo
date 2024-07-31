package bit

// 318. 最大单词长度乘积
// LCR 005. 最大单词长度乘积
// 给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
//
// 示例 1:
//
// 输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
// 输出: 16
// 解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
// 示例 2:
//
// 输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
// 输出: 4
// 解释: 这两个单词为 "ab", "cd"。
// 示例 3:
//
// 输入: words = ["a","aa","aaa","aaaa"]
// 输出: 0
// 解释: 不存在这样的两个单词。
//
// 提示：
//
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] 仅包含小写字母
//
// 注意：本题与主站 318 题相同：https://leetcode-cn.com/problems/maximum-product-of-word-lengths/
func maxProductWithLT(words []string) int {
	ans := 0
	masks := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a') // 按位或，aaccbb使用mask|=1<<(ch-'a')时，存储的信息只有abc，这样能节省空间，如果使用mask+=1<<(ch-'a')则会重复计算aabbcc，多计算一次，属实没必要
		}
		if len(word) > masks[mask] { // 当且仅当长度变长时更新，举例：aaabbb > aabb，这个时候更新长度
			masks[mask] = len(word)
		}
	}

	for x, lenX := range masks {
		for y, lenY := range masks {
			if x&y == 0 && lenX*lenY > ans {
				ans = lenX * lenY
			}
		}
	}
	return ans

}
func maxProduct(words []string) int {
	// 按照单词的长度对words数组排序
	//sort.Slice(words, func(i, j int) bool {
	//	return len(words[i]) > len(words[j])
	//})

	n := len(words)
	byteMap := make(map[string][26]int)
	for _, word := range words {
		tmp := [26]int{}
		for j := range word {
			tmp[word[j]-'a']++
		}
		byteMap[word] = tmp
	}
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isContains(byteMap[words[i]], byteMap[words[j]]) {
				continue
			} else {
				if len(words[i])*len(words[j]) > ans {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return ans
}

// a,b是否重叠
func isContains(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != 0 && b[i] != 0 {
			return true
		}
	}

	return false
}

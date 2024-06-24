package leetcode

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//提示：
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成

// 解题： 先遍历，找出最短的单词；然后遍历，以第一个单词为基准，取最短的单词长度，一个一个作比较，一旦比较出来都为0，出来就break
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	commonPrefix := strs[0][:minLen]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(commonPrefix); j++ {
			if commonPrefix[j] != strs[i][j] {
				minLen = j
				commonPrefix = commonPrefix[:j]
			}
		}
		if minLen == 0 {
			break
		}
	}

	return commonPrefix
}

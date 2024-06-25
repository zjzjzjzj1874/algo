package easy

import "sort"

// 49. 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:
// 输入: strs = [""]
// 输出: [[""]]
// 示例 3:
// 输入: strs = ["a"]
// 输出: [["a"]]
// 提示：
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母

// 解题：字母异位词，每个单词中字母出现的次数一定是相同的，随机位置无所谓了。所以可以把其字母排序后放入hash表中，hash的value是对应的[]string
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))

	strMap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sorted := string(s)

		val, ok := strMap[sorted]
		if !ok {
			sub := []string{str}
			strMap[sorted] = sub
		} else {
			val = append(val, str)
			strMap[sorted] = val
		}
	}

	for _, val := range strMap {
		result = append(result, val)
	}

	return result
}

// 解题：不用额外排序，直接使用单词个数，map的key按照数组来组织
func groupAnagramsV2(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	strMap := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, i := range str {
			key[i%26]++
		}
		val, ok := strMap[key]
		if !ok {
			sub := []string{str}
			strMap[key] = sub
		} else {
			val = append(val, str)
			strMap[key] = val
		}
	}

	for _, val := range strMap {
		result = append(result, val)
	}

	return result
}

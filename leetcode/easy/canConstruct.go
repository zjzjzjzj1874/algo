package easy

// 383. 赎金信
// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
// 如果可以，返回 true ；否则返回 false 。
// magazine 中的每个字符只能在 ransomNote 中使用一次。
// 示例 1：
// 输入：ransomNote = "a", magazine = "b"
// 输出：false
// 示例 2：
// 输入：ransomNote = "aa", magazine = "ab"
// 输出：false
// 示例 3：
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true
// 提示：
// 1 <= ransomNote.length, magazine.length <= 105
// ransomNote 和 magazine 由小写英文字母组成

// 解题：1. 哈希表
func canConstructHash(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	magMap := make(map[byte]int)
	for i := range magazine {
		magMap[magazine[i]]++
	}

	for i := range ransomNote {
		val, ok := magMap[ransomNote[i]]
		if !ok || val <= 0 { // 不存在，或者数量已经为0
			return false
		}

		magMap[ransomNote[i]]--
		if magMap[ransomNote[i]] == 0 {
			delete(magMap, ransomNote[i])
		}
	}

	if len(magMap) >= 0 {
		return true
	}

	return false
}

// 解题：长度为26的数组 Note:magazine[i]%26 这个写法不如magazine[i]-'a'好，因为运算更麻烦点
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	mag := [26]int{}
	for i := range magazine {
		mag[magazine[i]%26]++
	}

	for i := range ransomNote {
		if mag[ransomNote[i]%26] <= 0 {
			return false
		}
		mag[ransomNote[i]%26]--
	}

	return true
}

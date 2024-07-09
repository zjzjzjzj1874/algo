package easy

// 389. 找不同
// 给定两个字符串 s 和 t ，它们只包含小写字母。
//
// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
// 请找出在 t 中被添加的字母。
//
// 示例 1：
//
// 输入：s = "abcd", t = "abcde"
// 输出："e"
// 解释：'e' 是那个被添加的字母。
// 示例 2：
//
// 输入：s = "", t = "y"
// 输出："y"
//
// 提示：
//
// 0 <= s.length <= 1000
// t.length == s.length + 1
// s 和 t 只包含小写字母

// 解题：比较两个数组是否相同
func findTheDifference(s string, t string) byte {
	sm := [26]int{}
	st := [26]int{}

	for i := range s {
		sm[s[i]-'a']++
	}

	for i := range t {
		st[t[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		st[i] -= 1
		if st == sm {
			return byte(i + 'a')
		} else {
			st[i] += 1
		}
	}

	return byte(0)
}

// 解题：位运算
func findTheDifferenceWithBit(s string, t string) byte {
	ans := byte(0)

	for i := range s {
		ans = ans ^ s[i] ^ t[i]
	}

	return ans ^ t[len(t)-1]
}

// 解题：先++再--
func findTheDifferenceWithArr(s string, t string) byte {
	sm := [26]int{}

	for i := range s {
		sm[s[i]-'a']++
	}

	for i := range t {
		sm[t[i]-'a']--
		if sm[t[i]-'a'] < 0 {
			return t[i]
		}
	}

	return byte(0)
}

// 解题：求和
func findTheDifferenceWithSum(s string, t string) byte {
	sum := byte(0)

	for i := range t {
		sum += t[i]
	}

	for i := range s {
		sum -= s[i]
	}

	return sum
}

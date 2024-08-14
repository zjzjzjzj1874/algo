package doublePtr

// 541. 反转字符串 II
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
// 示例 1：
//
// 输入：s = "abcdefg", k = 2
// 输出："bacdfeg"
// 示例 2：
//
// 输入：s = "abcd", k = 2
// 输出："bacd"
//
// 提示：
//
// 1 <= s.length <= 104
// s 仅由小写英文组成
// 1 <= k <= 104
func reverseStrWithByte(s string, k int) string {
	ans := []byte(s)

	for i := 0; i < len(ans); i += 2 * k {
		sub := ans[i:min(i+k, len(s))]
		for l, r := 0, len(sub)-1; l <= r; {
			sub[l], sub[r] = sub[r], sub[l]
			l++
			r--
		}
	}
	return string(ans)
}
func reverseStr(s string, k int) string {
	count := 0
	n := len(s)
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		count++
		if count == 2*k {
			left := i + 1 - 2*k
			right := i + 1 - k - 1

			for left <= right {
				ans[left] = s[right]
				ans[right] = s[left]
				left++
				right--
			}
			for j := i + 1 - k; j <= i; j++ {
				ans[j] = s[j]
			}
			count = 0
		} else if i == n-1 {
			if count <= k { // 如果剩余字符少于 k 个，则将剩余字符全部反转。
				left := i + 1 - count
				right := i

				for left <= right {
					ans[left] = s[right]
					ans[right] = s[left]
					left++
					right--
				}
			} else {
				left := i + 1 - count
				right := left + k - 1

				for left <= right {
					ans[left] = s[right]
					ans[right] = s[left]
					left++
					right--
				}
				for j := i + 1 + k - count; j <= i; j++ {
					ans[j] = s[j]
				}
			}
		}
	}

	return string(ans)
}

// 344. 反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
//
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//
// 示例 1：
//
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
// 示例 2：
//
// 输入：s = ["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]
//
// 提示：
//
// 1 <= s.length <= 105
// s[i] 都是 ASCII 码表中的可打印字符
func reverseStringI(s []byte) {
	left := 0
	right := len(s) - 1

	for left <= right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}

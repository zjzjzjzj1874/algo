package math

// 415. 字符串相加
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
// 示例 1：
// 输入：num1 = "11", num2 = "123"
// 输出："134"
// 示例 2：
// 输入：num1 = "456", num2 = "77"
// 输出："533"
// 示例 3：
// 输入：num1 = "0", num2 = "0"
// 输出："0"
// 提示：
// 1 <= num1.length, num2.length <= 104
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	n := max(n1, n2)

	ans := make([]byte, n+1)
	shift := 0
	for i := 0; i < n; i++ {
		if i < n1 {
			shift += numByte[num1[n1-i-1]]
		}
		if i < n2 {
			shift += numByte[num2[n2-i-1]]
		}
		if shift >= 10 {
			ans[n-i] = byteNum[shift-10]
			shift /= 10
		} else {
			ans[n-i] = byteNum[shift]
			shift = 0
		}
	}
	if shift > 0 {
		ans[0] = byteNum[shift]
	} else {
		ans = ans[1:]
	}

	return string(ans)
}

var numByte = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

var byteNum = map[int]byte{
	0: '0',
	1: '1',
	2: '2',
	3: '3',
	4: '4',
	5: '5',
	6: '6',
	7: '7',
	8: '8',
	9: '9',
}

package math

// 204. 计数质数
// 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
// 示例 1：
//
// 输入：n = 10
// 输出：4
// 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
// 示例 2：
//
// 输入：n = 0
// 输出：0
// 示例 3：
//
// 输入：n = 1
// 输出：0
//
// 提示：
//
// 0 <= n <= 5 * 106
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	// 初始化，默认所有的数都是质数
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}

	ans := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	return ans
}

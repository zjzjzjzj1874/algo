package hwod

// 最多等和不相交连续子序列
// 题目描述
// 给定一个数组，我们称其中连续的元素为连续子序列，称这些元素的和为连续子序列的和。
// 数组中可能存在几组连续子序列，组内的连续子序列互不相交且有相同的和。
// 求一组连续子序列，组内子序列的数目最多。输出这个数目。
//
// 输入描述
// 第一行输入为数组长度𝑁，1≤N≤10^3。
// 第二行为N 个用空格分开的整数Ci: -10^5<=Ci<=10^5。
// 输出描述
// 第一行是一个整数M，表示满足要求的最多的组内子序列的数目。
// 示例一
// 输入
// 10
// 8 8 9 1 9 6 3 9 1 0
// 输出
// 4
// 说明
// 四个子序列第一个元素和最后一个元素下标分别为
// 2 2
// 4 4
// 5 6
// 7 7
// 示例一
// 输入
// 10
// -1 0 4 -3 6 5 -6 5 -7 -3
// 输出
// 3
// 说明
// 三个子序列第一个元素和最后一个元素下标分别为
// 3 3
// 5 8
// 9 9
func maxConSeq(seq []int) (ans int) {
	n := len(seq)
	dp := make([]int, n)
	copy(dp, seq)
	sumCount := make(map[int]int)
	sumPos := make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j+i < n; j++ {
			if i > 0 {
				dp[j] = dp[j] + seq[j+i]
			}
			sum := dp[j]
			if _, exists := sumCount[sum]; !exists {
				sumCount[sum] = 0
				sumPos[sum] = make(map[int]struct{})
			}

			exists := false
			poss := sumPos[sum]
			for k := j; k <= j+i; k++ {
				if _, found := poss[k]; found {
					exists = true
					break
				}
			}
			if !exists {
				newSum := sumCount[sum] + 1
				sumCount[sum] = newSum
				ans = max(ans, newSum)
				for k := j; k <= j+i; k++ {
					poss[k] = struct{}{}
				}
				sumPos[sum] = poss
			}
		}
	}

	return
}

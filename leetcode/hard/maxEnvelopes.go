package hard

import "sort"

// 354. 俄罗斯套娃信封问题
// 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
// 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出：3
// 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
// 示例 2：
//
// 输入：envelopes = [[1,1],[1,1],[1,1]]
// 输出：1
//
// 提示：
//
// 1 <= envelopes.length <= 105
// envelopes[i].length == 2
// 1 <= wi, hi <= 105

// maxEnvelopes 计算最多能嵌套的信封数
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	// 按宽度升序排列，如果宽度相同，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	// 备忘录，用于存储子问题的解
	memo := make([]int, n)

	// 递归函数，计算从第 i 个信封开始的最大嵌套数
	var dp func(int) int
	dp = func(i int) int {
		if memo[i] != 0 {
			return memo[i]
		}
		maxEnvelopes := 1 // 至少包含当前信封
		for j := i + 1; j < n; j++ {
			if envelopes[j][0] > envelopes[i][0] && envelopes[j][1] > envelopes[i][1] {
				maxEnvelopes = max(maxEnvelopes, 1+dp(j))
			}
		}
		memo[i] = maxEnvelopes
		return maxEnvelopes
	}

	// 计算所有信封的最大嵌套数
	result := 0
	for i := 0; i < n; i++ {
		result = max(result, dp(i))
	}
	return result
}

package matrix

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
func generateMatrix(n int) (ans [][]int) {
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	start := 1

	for left <= right && top <= bottom {
		for i := top; i <= right; i++ {
			ans[top][i] = start
			start++
		}
		top++
		for i := top; i <= bottom; i++ {
			ans[i][right] = start
			start++
		}
		right--
		for i := right; i >= left; i-- {
			ans[bottom][i] = start
			start++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			ans[i][left] = start
			start++
		}
		left++
	}

	return
}

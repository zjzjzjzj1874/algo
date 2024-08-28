package hwod

import "math"

// 找数字 TODO 待完成
// 题目描述
// 给一个二维数组 nums，对于每一个元素 num[i]，找出距离最近的且值相等的元素，输出横 纵坐标差值的绝对值之和，如果没有等值元素，则输出-1。
// 例如:
// 输入数组 nums 为
// 03542
// 25783
// 25424
// 对于 num[0][0] = 0，不存在相等的值。
// 对于 num[0][1] = 3，存在一个相等的值，最近的坐标为 num[1][4]，最小距离为 4。 对于 num[0][2] = 5，存在两个相等的值，最近的坐标为 num[1][1]，故最小距离为 2。 ...
// 对于 num[1][1] = 5，存在两个相等的值，最近的坐标为 num[2][1]，故最小距离为 1。 ...
// 故输出为
// -14233
// 11-1-14 11232
// 输入描述
// 输入第一行为二维数组的行
// 输入第二行为二维数组的列
// 输入的数字以空格隔开。
func findNum(nums [][]int) (target [][]int) {
	m := len(nums)
	n := len(nums[0])

	target = make([][]int, m)
	for i := range target {
		target[i] = make([]int, n)
		for j := range target[i] {
			target[i][j] = -1
		}
	}
	// [x,y]代表num，i，j表示需要确认的数的坐标
	var dfs func(x, y, i, j int) int
	dfs = func(x, y, i, j int) int {
		if i >= m || j >= n || j < 0 {
			return -1
		}
		tar := -1
		if nums[i][j] == nums[x][y] && (i != x || j != y) {
			tar = int(math.Abs(float64(i-x)) + math.Abs(float64(j-y)))
			target[i][j] = tar
			return tar
		}
		// 向右寻找
		if tar = dfs(x, y, i, j+1); tar != -1 {
			target[i][j+1] = tar
			return tar
		}
		// 向下寻找
		if tar = dfs(x, y, i+1, j); tar != -1 {
			target[i+1][j] = tar
			return tar
		}
		// 向右下寻找
		if tar = dfs(x, y, i+1, j+1); tar != -1 {
			target[i+1][j+1] = tar
			return tar
		}
		// 向左下寻找
		if tar = dfs(x, y, i+1, j-1); tar != -1 {
			target[i+1][j-1] = tar
			return tar
		}
		return tar
	}
	for i := range nums {
		for j := range nums[i] {
			tar := dfs(i, j, i, j)
			if target[i][j] != -1 && tar < target[i][j] {
				target[i][j] = tar
			}
		}
	}

	return
}

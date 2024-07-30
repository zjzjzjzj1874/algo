package prefixsum

// LCR 013. 二维区域和检索 - 矩阵不可变
// 给定一个二维矩阵 matrix，以下类型的多个请求：
//
// 计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
// 实现 NumMatrix 类：
//
// NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
// int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
//
// 示例 1：
//
// 输入:
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
// [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
// 输出:
// [null, 8, 11, 12]
//
// 解释:
// NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// -105 <= matrix[i][j] <= 105
// 0 <= row1 <= row2 < m
// 0 <= col1 <= col2 < n
// 最多调用 104 次 sumRegion 方法
//
// 注意：本题与主站 304 题相同： https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	matrix    [][]int // 原始矩阵
	SumMatrix [][]int // 矩阵前缀和
}

func ConstructorRegion(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := len(matrix[0])
	nm := NumMatrix{
		matrix: matrix,
	}
	sm := make([][]int, n)
	for i := range matrix {
		sm[i] = make([]int, m)
		sum := 0
		for j := 0; j < m; j++ {
			sum += matrix[i][j]
			sm[i][j] = sum
		}
	}
	for j := 0; j < m; j++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum += sm[i][j]
			sm[i][j] = sum
		}
	}

	nm.SumMatrix = sm

	return nm
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := n.SumMatrix[row2][col2]
	if row1 != 0 {
		ans -= n.SumMatrix[row1-1][col2]
	}
	if col1 != 0 {
		ans -= n.SumMatrix[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		ans += n.SumMatrix[row1-1][col1-1]
	}
	return ans
}

// 求前缀和，index1到index2的和 = sum2 - sum1

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

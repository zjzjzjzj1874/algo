package graph

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1：
//
// 输入：grid = [
// ["1","1","1","1","0"],
// ["1","1","0","1","0"],
// ["1","1","0","0","0"],
// ["0","0","0","0","0"]
// ]
// 输出：1
// 示例 2：
//
// 输入：grid = [
// ["1","1","0","0","0"],
// ["1","1","0","0","0"],
// ["0","0","1","0","0"],
// ["0","0","0","1","1"]
// ]
// 输出：3
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'

// 解题：递归感染
func numIslands(grid [][]byte) int {
	ans := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				infect(grid, i, j, m, n)
			}
		}
	}

	return ans
}

// up,down表示矩阵的(高)，不超过m
// left,right表示矩阵的宽，不超过n
// i,j
func infect(grid [][]byte, i, j, m, n int) {
	if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	infect(grid, i-1, j, m, n)
	infect(grid, i+1, j, m, n)
	infect(grid, i, j-1, m, n)
	infect(grid, i, j+1, m, n)
}

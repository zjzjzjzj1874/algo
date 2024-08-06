package math

// 840. 矩阵中的幻方
// 3 x 3 的幻方是一个填充有 从 1 到 9  的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。
//
// 给定一个由整数组成的row x col 的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
//
// 示例 1：
//
// 输入: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]
// 输出: 1
// 解释:
// 下面的子矩阵是一个 3 x 3 的幻方：
//
// 而这一个不是：
//
// 总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
// 示例 2:
//
// 输入: grid = [[8]]
// 输出: 0
//
// 提示:
//
// row == grid.length
// col == grid[i].length
// 1 <= row, col <= 10
// 0 <= grid[i][j] <= 15

// 解题：暴力尝试
func numMagicSquaresInside(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if n < 3 || m < 3 {
		return 0
	}

	z_n := [16]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	ans := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if grid[i+1][j+1] != 5 {
				continue // 过滤
			}
			tmp := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			c := [16]int{}
			if grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] != tmp {
				continue
			}
			if grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] != tmp {
				continue
			}
			if grid[i][j]+grid[i+1][j]+grid[i+2][j] != tmp {
				continue
			}
			if grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] != tmp {
				continue
			}
			if grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] != tmp {
				continue
			}
			if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != tmp {
				continue
			}
			if grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] != tmp {
				continue
			}
			c[grid[i][j]]++
			c[grid[i][j+1]]++
			c[grid[i][j+2]]++

			c[grid[i+1][j]]++
			c[grid[i+1][j+1]]++
			c[grid[i+1][j+2]]++

			c[grid[i+2][j]]++
			c[grid[i+2][j+1]]++
			c[grid[i+2][j+2]]++

			if c != z_n {
				continue
			}
			ans++
		}
	}

	return ans
}

package middle

import "fmt"

// 289. 生命游戏
// 根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
//
// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
//
// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。
//
// 示例 1：
//
// 输入：board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// 输出：[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
// 示例 2：
//
// 输入：board = [[1,1],[1,0]]
// 输出：[[1,1],[1,1]]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 25
// board[i][j] 为 0 或 1
//
// 进阶：
//
// 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
// 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

// 解题：使用额外的空间。记录原始的值
func gameOfLifeWithOMN(board [][]int) {
	m := len(board)
	n := len(board[0])

	ans := make([][]int, len(board))
	for i := range ans {
		ans[i] = make([]int, len(board[i]))
	}

	var calOne func(i, j int) int
	calOne = func(i, j int) (alive int) {
		up := i - 1
		down := i + 1
		l := j - 1
		r := j + 1
		if l >= 0 {
			if board[i][l] == 1 {
				alive++
			}
			if up >= 0 && board[up][l] == 1 {
				alive++
			}
			if down < m && board[down][l] == 1 {
				alive++
			}
		}
		if r < n {
			if board[i][r] == 1 {
				alive++
			}
			if up >= 0 && board[up][r] == 1 {
				alive++
			}
			if down < m && board[down][r] == 1 {
				alive++
			}
		}
		if up >= 0 && board[up][j] == 1 {
			alive++
		}
		if down < m && board[down][j] == 1 {
			alive++
		}

		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 1 && j == 0 {
				fmt.Println("注意")
			}
			ans[i][j] = board[i][j]
			alive := calOne(i, j)
			if alive < 2 || alive > 3 {
				ans[i][j] = 0
			}
			if alive == 3 {
				ans[i][j] = 1
			}
		}
	}

	copy(board, ans)
}

// 思路2：使用复合状态
// 规则 1：如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡。这时候，将细胞值改为 -1，代表这个细胞过去是活的现在死了；
// 规则 2：如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活。这时候不改变细胞的值，仍为 1；
// 规则 3：如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡。这时候，将细胞的值改为 -1，代表这个细胞过去是活的现在死了。可以看到，因为规则 1 和规则 3 下细胞的起始终止状态是一致的，因此它们的复合状态也一致；
// 规则 4：如果死细胞周围正好有三个活细胞，则该位置死细胞复活。这时候，将细胞的值改为 2，代表这个细胞过去是死的现在活了。
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	var calNeibor func(i, j int) int // 计算邻居细胞的存活数量
	calNeibor = func(i, j int) (ans int) {
		for x := i - 1; x <= i+1; x++ {
			for y := j - 1; y <= j+1; y++ {
				if x == i && y == j { // 不能重复计数
					continue
				}
				if (x >= 0 && x < m) && (y < n && y >= 0) {
					if board[x][y] == 1 || board[x][y] == -1 { // 活着，或者曾经活着
						ans++
					}
				}
			}
		}

		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			alive := calNeibor(i, j)
			// 规则1，3
			if (alive < 2 || alive > 3) && (board[i][j] == 1) {
				board[i][j] = -1
			}

			// 规则四：曾经死亡，现在活了
			if alive == 3 && board[i][j] == 0 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] <= 0 {
				board[i][j] = 0
			} else {
				board[i][j] = 1
			}
		}
	}

	fmt.Println(board)
}

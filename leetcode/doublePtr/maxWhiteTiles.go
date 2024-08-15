package doublePtr

import "sort"

// 2271. 毯子覆盖的最多白色砖块数
// 给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。
//
// 同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子的长度。
//
// 请你返回使用这块毯子，最多 可以盖住多少块瓷砖。
//
// 示例 1：
//
// 输入：tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
// 输出：9
// 解释：将毯子从瓷砖 10 开始放置。
// 总共覆盖 9 块瓷砖，所以返回 9 。
// 注意可能有其他方案也可以覆盖 9 块瓷砖。
// 可以看出，瓷砖无法覆盖超过 9 块瓷砖。
// 示例 2：
//
// 输入：tiles = [[10,11],[1,1]], carpetLen = 2
// 输出：2
// 解释：将毯子从瓷砖 10 开始放置。
// 总共覆盖 2 块瓷砖，所以我们返回 2 。
//
// 提示：
//
// 1 <= tiles.length <= 5 * 104
// tiles[i].length == 2
// 1 <= li <= ri <= 109
// 1 <= carpetLen <= 109
// tiles 互相 不会重叠 。

// 这个解题会超过时间限制，推荐用下面的
func maximumWhiteTilesTimeout(tiles [][]int, carpetLen int) int {
	n := len(tiles)
	if n == 1 {
		return min(carpetLen, tiles[0][1]-tiles[0][0]+1)
	}

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	ans := 0

	l := 0
	for ; l < n; l++ {
		title := tiles[l]
		// 当前长度能覆盖carpetLen，取较大值
		if title[1]-title[0]+1 >= carpetLen {
			ans = max(ans, title[1]-title[0]+1)
			continue
		}
		tmp := title[1] - title[0] + 1
		r := l + 1
		for ; r < n && tiles[r][1]-tiles[l][0]+1 < carpetLen; r++ {
			tmp += tiles[r][1] - tiles[r][0] + 1
		}
		// 这里处理部分覆盖的情况
		if r < n && tiles[r][0] <= tiles[l][0]+carpetLen-1 {
			ans = max(ans, tmp+tiles[l][0]+carpetLen-tiles[r][0])
		} else {
			ans = max(ans, tmp)
		}
	}

	return ans
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	// 首先对地砖区间进行排序
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	maxCovered := 0
	currentCovered := 0
	right := 0
	n := len(tiles)

	for left := 0; left < n; left++ {
		// 移动右指针使得当前地毯能够覆盖尽可能多的地砖
		for right < n && tiles[right][1]-tiles[left][0]+1 <= carpetLen {
			currentCovered += tiles[right][1] - tiles[right][0] + 1
			right++
		}

		if right < n && tiles[right][0] <= tiles[left][0]+carpetLen-1 {
			// 处理部分覆盖的情况
			maxCovered = max(maxCovered, currentCovered+(tiles[left][0]+carpetLen-tiles[right][0]))
		} else {
			maxCovered = max(maxCovered, currentCovered)
		}

		// 移动左指针，并移除对应区间的覆盖面积
		currentCovered -= tiles[left][1] - tiles[left][0] + 1
	}

	return maxCovered
}

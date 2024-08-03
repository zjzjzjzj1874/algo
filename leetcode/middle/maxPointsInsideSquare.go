package middle

// 3143. 正方形中的最多点数
// 给你一个二维数组 points 和一个字符串 s ，其中 points[i] 表示第 i 个点的坐标，s[i] 表示第 i 个点的 标签 。
//
// 如果一个正方形的中心在 (0, 0) ，所有边都平行于坐标轴，且正方形内 不 存在标签相同的两个点，那么我们称这个正方形是 合法 的。
//
// 请你返回 合法 正方形中可以包含的 最多 点数。
//
// 注意：
//
// 如果一个点位于正方形的边上或者在边以内，则认为该点位于正方形内。
// 正方形的边长可以为零。
//
// 示例 1：
//
// 输入：points = [[2,2],[-1,-2],[-4,4],[-3,1],[3,-3]], s = "abdca"
//
// 输出：2
//
// 解释：
//
// 边长为 4 的正方形包含两个点 points[0] 和 points[1] 。
//
// 示例 2：
//
// 输入：points = [[1,1],[-2,-2],[-2,2]], s = "abb"
//
// 输出：1
//
// 解释：
//
// 边长为 2 的正方形包含 1 个点 points[0] 。
//
// 示例 3：
//
// 输入：points = [[1,1],[-1,-1],[2,-2]], s = "ccd"
//
// 输出：0
//
// 解释：
//
// 任何正方形都无法只包含 points[0] 和 points[1] 中的一个点，所以合法正方形中都不包含任何点。
//
// 提示：
//
// 1 <= s.length, points.length <= 105
// points[i].length == 2
// -109 <= points[i][0], points[i][1] <= 109
// s.length == points.length
// points 中的点坐标互不相同。
// s 只包含小写英文字母。

// 解题思路：对于一个点 (x,y)，我们可以将其映射到以原点为中心的第一象限，即 (max(∣x∣,∣y∣),max(∣x∣,∣y∣))。
// 这样，我们就可以将所有的点映射到第一象限，然后按照点到原点的距离进行排序。
// 我们可以使用哈希表 g 来存储所有点到原点的距离，然后按照距离进行排序。对于每个距离 d，我们将所有距离为 d 的点放在一起，然后遍历这些点，
// 如果有两个点的标签相同，那么这个正方形是不合法的，直接返回答案。否则，我们将这些点加入到答案中。
func maxPointsInsideSquare(points [][]int, s string) (ans int) {

	// 对points按照从大到小排序
	//sort.Slice(points, func(i, j int) bool {
	//	imax := max(math.Abs(float64(points[i][0])), math.Abs(float64(points[i][1])))
	//	jmax := max(math.Abs(float64(points[j][0])), math.Abs(float64(points[j][1])))
	//
	//	return imax > jmax
	//})
	//
	//tagMap := make(map[[2]int]byte)
	//for i := range points {
	//	tagMap[[2]int(points[i])] = s[i]
	//}
	//n := len(points)
	//
	//for i, point := range points {
	//	//maxLen := max(points[i][0], points[i][1])
	//	tMap := make(map[byte]bool)
	//	tmpAns := 1 // 包含当前点
	//
	//	tMap[tagMap[[2]int(point)]] = true
	//	//tMap[s[i]] = true
	//	for j := i + 1; j < n; j++ {
	//		if tMap[s[j]] {
	//			tmpAns = 0
	//			break
	//		}
	//		tMap[s[j]] = true
	//		tmpAns++
	//	}
	//	if tmpAns > ans {
	//		ans = tmpAns
	//	}
	//}

	return
}

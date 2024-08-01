package middle

import "sort"

// LCP 40. 心算挑战
// 「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。 给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
//
// 示例 1：
//
// 输入：cards = [1,2,8,9], cnt = 3
//
// 输出：18
//
// 解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。
//
// 示例 2：
//
// 输入：cards = [3,3,1], cnt = 1
//
// 输出：0
//
// 解释：不存在获取有效得分的卡牌方案。
//
// 提示：
//
// 1 <= cnt <= cards.length <= 10^5
// 1 <= cards[i] <= 1000

// 数据规模太大，回溯一定会超时。考虑其他的贪心算法
func maxmiumScore(cards []int, cnt int) (ans int) {
	// 从大到小排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	tmp := 0            // 临时变量，存sum
	odd, even := -1, -1 // 存最接近的奇偶数
	for i := 0; i < cnt; i++ {
		tmp += cards[i]
		if cards[i]&1 == 0 { // 偶数
			even = cards[i]
		} else {
			odd = cards[i]
		}
	}
	if tmp&1 == 0 {
		return tmp
	}
	// tmp是一个奇数
	for i := cnt; i < len(cards); i++ {
		if cards[i]&1 != 0 { // 当前是个奇数，拿走最小的偶数，加上当前这个奇数
			if even != -1 {
				ans = max(ans, tmp-even+cards[i])
			}
		} else { // 当前数是偶数，去掉最近(最小)的奇数，用这个偶数
			if odd != -1 {
				ans = max(ans, tmp-odd+cards[i])
			}
		}
	}

	return
}

func maxmiumScoreTimeout(cards []int, cnt int) (ans int) {
	choice := make([]int, 0, cnt)
	n := len(cards)
	selected := make([]bool, n)
	sum := 0
	var backtrack func(i int) // i表示下一个带选择的选项

	backtrack = func(i int) {
		if len(choice) == cnt {
			if sum%2 == 0 && sum > ans {
				ans = sum
			}
			return
		}

		for j := i; j < n; j++ {
			if selected[j] {
				continue // 选择过，不能重复选择
			}

			selected[j] = true
			choice = append(choice, cards[j])
			sum += cards[j]
			backtrack(j + 1)
			// 恢复，不选择这个
			choice = choice[:len(choice)-1]
			sum -= cards[j]
			selected[j] = false
		}
	}

	backtrack(0)
	return
}

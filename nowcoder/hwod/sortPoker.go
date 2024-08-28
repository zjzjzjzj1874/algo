package hwod

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 整理扑克牌:https://wiki.amoscloud.com/zh/ProgramingPractice/NowCoder/Adecco/Topic0198
// 题目描述
// 给定一组数字，表示扑克牌的牌面数字，忽略扑克牌的花色，请按如下规则对这一组扑克 牌进行整理:
// 步骤 1、对扑克牌进行分组，形成组合牌，规则如下:
// 1. 当牌面数字相同张数 大于等于 4 时 ，组合牌为“炸弹”;
// 2. 3 张相同牌面数字 +
// 2 张相同牌面数字，且 3 张牌与 2 张牌不相同时，组合牌为“葫芦”;
// 3. 3 张相同牌面数字，组合牌为“三张”;
// 4. 2 张相同牌面数字，组合牌为“对子”;
// 5. 剩余没有相同的牌，则为“单张”;
// 步骤 2、对上述组合牌进行由大到小排列，规则如下:
// 1. 不同类型组合牌之间由大到小排列规则:“炸弹” > "葫芦" > "三张" > "对子" > “单张”;
// 2. 相同类型组合牌之间，除“葫芦”外，按组合牌全部牌面数字加总由大到小排列;
// 3. “葫芦”则先按 3 张相同牌面数字加总由大到小排列，3 张相同牌面数字加总相同时， 再按另外 2 张牌面数字加总由大到小排列;
// 4. 由于“葫芦”
// >“三张”，因此如果能形成更大的组合牌，也可以将“三张”拆分为 2 张和 1 张，其中
// 的 2 张可以和其它“三张”重新组合成“葫芦”，剩下的 1 张为“单张”
// 步骤 3、当存在多个可能组合方案时，按如下规则排序取最大的一个组合方案:
// 1. 依次对组合方案中的组合牌进行大小比较，规则同上;
// 2. 当组合方案 A 中的第 n 个组合牌大于组合方案 B 中的第 n 个组合牌时，组合方案 A 大于 组合方案 B;
// 输入描述
// 第一行为空格分隔的 N 个正整数，每个整数取值范围[1,13]，N 的取值范围[1,1000]
// 输出描述 经重新排列后的扑克牌数字列表，每个数字以空格分隔

type Card struct {
	Value int
	Count int
}

func sortPoker(poker string) {
	solution(poker)
}

func solution(line string) {
	// Parse input
	tokens := strings.Fields(line)
	cardMap := make(map[int]int)
	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		cardMap[num]++
	}

	// Convert map to a slice of Card
	cards := []Card{}
	for value, count := range cardMap {
		cards = append(cards, Card{Value: value, Count: count})
	}

	// Sort cards by count (descending) and then by value (descending)
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Count != cards[j].Count {
			return cards[i].Count > cards[j].Count
		}
		return cards[i].Value > cards[j].Value
	})

	var builder strings.Builder

	for len(cards) > 0 {
		// Sort cards each iteration
		sort.Slice(cards, func(i, j int) bool {
			if cards[i].Count != cards[j].Count {
				return cards[i].Count > cards[j].Count
			}
			return cards[i].Value > cards[j].Value
		})

		// Check for Bomb
		if cards[0].Count >= 4 {
			appendCards(&builder, cards[0].Value, 4)
			if cards[0].Count == 4 {
				cards = cards[1:]
			} else {
				cards[0].Count -= 4
			}
			continue
		}

		// Check for FullHouse or Three of a Kind
		if cards[0].Count == 3 && len(cards) > 1 && cards[1].Count >= 2 {
			appendCards(&builder, cards[0].Value, 3)
			appendCards(&builder, cards[1].Value, 2)
			if cards[1].Count == 2 {
				cards = append(cards[:1], cards[2:]...)
			} else {
				cards[1].Count -= 2
			}
			cards = cards[1:]
			continue
		} else if cards[0].Count == 3 {
			appendCards(&builder, cards[0].Value, 3)
			cards = cards[1:]
			continue
		}

		// Process remaining cards
		for i := 0; i < len(cards); i++ {
			appendCards(&builder, cards[i].Value, cards[i].Count)
		}
		cards = nil
	}

	fmt.Print(strings.TrimSpace(builder.String()))
}

func appendCards(builder *strings.Builder, number, count int) {
	for i := 0; i < count; i++ {
		builder.WriteString(fmt.Sprintf("%d ", number))
	}
}

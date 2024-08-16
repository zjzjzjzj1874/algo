package greedy

import "sort"

// 1647. 字符频次唯一的最小删除次数
// 如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。
//
// 给你一个字符串 s，返回使 s 成为 优质字符串 需要删除的 最小 字符数。
//
// 字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，'a' 的频次是 2，而 'b' 的频次是 1 。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：0
// 解释：s 已经是优质字符串。
// 示例 2：
//
// 输入：s = "aaabbbcc"
// 输出：2
// 解释：可以删除两个 'b' , 得到优质字符串 "aaabcc" 。
// 另一种方式是删除一个 'b' 和一个 'c' ，得到优质字符串 "aaabbc" 。
// 示例 3：
//
// 输入：s = "ceabaacb"
// 输出：2
// 解释：可以删除两个 'c' 得到优质字符串 "eabaab" 。
// 注意，只需要关注结果字符串中仍然存在的字符。（即，频次为 0 的字符会忽略不计。）
//
// 提示：
//
// 1 <= s.length <= 105
// s 仅含小写英文字母
func minDeletions(s string) (ans int) {
	alpha := [26]int{}

	for _, char := range s {
		alpha[char-'a']++
	}
	nums := make([]int, 0, 8)
	for _, v := range alpha {
		if v != 0 {
			nums = append(nums, v)
		}
	}

	sort.Ints(nums)

	countMap := make(map[int]struct{})
	for i := len(nums) - 1; i >= 0; i-- {
		_, ok := countMap[nums[i]]
		if !ok {
			countMap[nums[i]] = struct{}{}
			continue
		}
		// 存在过，依次删除
		nums[i] = nums[i] - 1
		ans++
		for {
			_, ok := countMap[nums[i]]
			if ok && nums[i] != 0 {
				ans++
				nums[i] = nums[i] - 1
			} else {
				countMap[nums[i]] = struct{}{}
				break
			}
		}
	}

	return
}

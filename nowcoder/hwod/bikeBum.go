package hwod

import "sort"

// 租车骑绿道
// 题目描述
// 部门组织绿道骑行团建活动。租用公共双人自行车骑行，每辆自行车最多坐两人、做大载 重 M。
// 给出部门每个人的体重，请问最多需要租用多少双人自行车。
// 输入描述
// 第一行两个数字 m、n，自行车限重 m，代表部门总人数 n。
// 第二行，n 个数字，代表每个人的体重。体重都小于等于自行车限重 m。 0<m<=200
// 0<n<=1000000
// 输出描述 最小需要的双人自行车数量。
func minBikeNum(m, n int, weight []int) (ans int) {
	sort.Ints(weight) // 部门人按照体重由小到大排序

	l := 0
	r := n - 1
	for l <= r {
		// 体重<=自行车重量，所以优化：
		if weight[l]+weight[r] <= m {
			l++
		}
		r--
		ans++
	}

	return
}

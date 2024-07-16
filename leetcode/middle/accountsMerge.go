package middle

import "sort"

// 721. 账户合并
// 给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。
//
// 现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。
//
// 合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。
// 示例 1：
//
// 输入：accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
// 输出：[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
// 解释：
// 第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。
// 第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
// 可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
// ['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。
// 示例 2：
//
// 输入：accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]
// 输出：[["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]
//
// 提示：
//
// 1 <= accounts.length <= 1000
// 2 <= accounts[i].length <= 10
// 1 <= accounts[i][j].length <= 30
// accounts[i][0] 由英文字母组成
// accounts[i][j] (for j > 0) 是有效的邮箱地址

// 解题：暴力解法，并查集
func accountsMerge(accounts [][]string) [][]string {
	aMap := make(map[string][][]string) // map[Name][][]string

	for i := range accounts {
		account := accounts[i]
		val, ok := aMap[account[0]]
		if !ok { // 不存在
			aMap[account[0]] = [][]string{account}
			continue
		}

		// 已经存在，判断是否是同一个人
		sameIdx := -1
		var sameEmail []string
		for j := range val { // 遍历每一个人
			if isSame, emails := isSameAccount(account, val[j]); isSame {
				sameIdx = j
				sameEmail = emails
			}
		}
		if sameIdx != -1 {
			// 是同一个人
			account = []string{account[0]}
			account = append(account, sameEmail...)
			val[sameIdx] = account
		} else {
			val = append(val, account)
		}
		aMap[account[0]] = val
	}

	ans := make([][]string, 0, len(aMap))
	for _, i := range aMap {
		for j := range i {
			na := duplicateEmail(i[j])
			sort.Strings(na[1:])
			ans = append(ans, na)
		}
	}

	return ans
}

// 判断是否是同一个人，是同一个人，对其email进行合并
func isSameAccount(a1, a2 []string) (isSame bool, combineEmails []string) {
	if a1[0] != a2[0] {
		return
	}

	aMap := make(map[string]struct{}) // map[email]struct{}
	for i := 1; i < len(a1); i++ {
		aMap[a1[i]] = struct{}{}
	}

	for i := 1; i < len(a2); i++ {
		if _, ok := aMap[a2[i]]; ok {
			isSame = true
			break
		}
	}
	if isSame {
		for i := 1; i < len(a2); i++ {
			if _, ok := aMap[a2[i]]; !ok {
				aMap[a2[i]] = struct{}{}
			}
		}
	} else {
		return
	}

	for email := range aMap {
		combineEmails = append(combineEmails, email)
	}
	return
}

// 对邮箱去重
func duplicateEmail(account []string) (newAccount []string) {
	aMap := make(map[string]struct{}) // map[email]struct{}
	for i := 1; i < len(account); i++ {
		aMap[account[i]] = struct{}{}
	}

	newAccount = append(newAccount, account[0])
	for email := range aMap {
		newAccount = append(newAccount, email)
	}

	return
}

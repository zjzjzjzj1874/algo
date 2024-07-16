package leetcode

import "fmt"

// KMP算法: 判断str2是否是str1的子串

func KMP(str1, str2 string) int {
	m := len(str1)
	n := len(str2)

	i1 := 0
	i2 := 0
	next := getNext(str2)
	for i1 < m && i2 < n {
		if str1[i1] == str2[i2] {
			i1++
			i2++
		} else if i2 == 0 {
			i1++
		} else {
			i2 = next[i2-1]
		}
	}

	if i2 == n {
		return i1 - i2
	}

	return -1
}

func getNext(str2 string) []int {
	if len(str2) == 0 {
		return []int{}
	}
	if len(str2) == 1 {
		return []int{-1}
	}

	next := make([]int, len(str2))
	next[0] = 0

	cn := 0
	for i := 1; i < len(str2); {
		if str2[i] == str2[cn] {
			cn++
			next[i] = cn
			i++
		} else if cn > 0 {
			cn = next[cn-1]
		} else {
			next[i] = 0
			i++
		}
	}

	return next
}

// GetNext 减1
func GetNext(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{-1}
	}
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	// next[n-1]+1 最长相同前后缀的长度

	fmt.Println(next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0)
	fmt.Println(next[n-1] != -1)
	fmt.Println(n % (n - (next[n-1] + 1)))

	return next
}

// GetNextV1 不减0
func GetNextV1(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{-1}
	}

	next := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	// next[n-1]+1 最长相同前后缀的长度

	fmt.Println(next[n-1] != 0 && n%(n-(next[n-1])) == 0)
	fmt.Println(next[n-1] != -1)
	fmt.Println(n % (n - (next[n-1] + 1)))

	return next
}

// GetNextV2 不减0
func GetNextV2(s string) []int {
	lps := make([]int, len(s))
	n := len(s)
	length := 0
	i := 1

	for i < len(s) {
		if s[i] == s[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	fmt.Println(lps[n-1] != 0 && n%(n-(lps[n-1])) == 0)
	fmt.Println(lps[n-1] != -1)
	fmt.Println(n % (n - (lps[n-1] + 1)))

	return lps
}

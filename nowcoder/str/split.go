package main

import "fmt"

// HJ4 字符串分隔
// 描述
// •输入一个字符串，请按长度为8拆分每个输入字符串并进行输出；
//
// •长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
// 输入描述：
// 连续输入字符串(每个字符串长度小于等于100)
//
// 输出描述：
// 依次输出所有分割后的长度为8的新字符串
//
// 示例1
// 输入：
// abc
// 复制
// 输出：
// abc00000
func split() {
	for {
		str := ""
		fmt.Scan(&str)
		if str == "" {
			continue
		}

		n := len(str)
		ans := ""
		for i := 0; i < n; i++ {
			if i == n-1 {
				ans += str[n-1:]
			} else {
				ans += str[i : i+1]
			}
			if (i+1)%8 == 0 {
				fmt.Println(ans)
				ans = ""
			}
		}
		for len(ans) < 8 {
			ans += "0"
		}
		fmt.Println(ans)
	}
}

// 解题：双指针-滑动窗口
func slidingWindow() {
	//func main() {
	for {
		str := ""
		fmt.Scan(&str)
		if str == "" {
			continue
		}

		n := len(str)
		l := 0
		r := 0
		for ; r < n; r++ {
			if r-l+1 == 8 {
				fmt.Println(str[l : r+1])
				l = r + 1
			}
		}
		if r == l {
			continue
		}
		for r-l < 8 {
			str += "0"
			r++
		}
		fmt.Println(str[l:r])
	}
}

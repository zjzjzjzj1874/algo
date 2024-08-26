package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 堆排的实现 --> 堆排不稳定，不是该题的解决方案
// HJ68 成绩排序

func main() {
	n := 0    // 数量
	sort := 0 // 排序
	fmt.Scan(&n)
	fmt.Scan(&sort)

	sc := bufio.NewReader(os.Stdin)
	stus := make([]student, n)
	for i := 0; i < n; i++ {
		msg, _ := sc.ReadString('\n')
		msgs := strings.Fields(msg)
		grade, _ := strconv.Atoi(msgs[1])
		stus[i] = student{
			name:  msgs[0],
			grade: grade,
		}
	}

	qhelper(stus, 0, n-1)
	for i := range stus {
		if sort == 1 {
			fmt.Println(stus[i].name, stus[i].grade)
		} else {
			fmt.Println(stus[n-1-i].name, stus[n-1-i].grade)
		}
	}
}

func qhelper(stu []student, l, r int) {
	if l < r {
		region := partitionV2(stu, l, r)

		qhelper(stu, l, region[0]-1)
		qhelper(stu, region[1]+1, r)
	}
}

// partition，类似荷兰国旗
//func partition(stu []student, l, r int) []int {
//	less := l - 1
//	more := r + 1
//	base := stu[r] // 基准值
//	for l < more {
//		if stu[l].grade < base.grade { // 左边小于基准值，交换位置
//			stu[l], stu[less+1] = stu[less+1], stu[l]
//			l++
//			less++
//		} else if stu[l].grade == base.grade { // 相等,直接移动指针
//			l++
//		} else { // 大于，把左边和最后元素交换位置，然后移动最右边元素指针，接着l继续和基准值比较
//			stu[l], stu[more-1] = stu[more-1], stu[l]
//			more--
//		}
//	}
//
//	return []int{less + 1, more - 1}
//}

func partitionV2(stu []student, l, r int) []int {
	less := l - 1
	more := r + 1
	base := stu[r].grade

	for l < more {
		if stu[l].grade < base {
			stu[l], stu[less+1] = stu[less+1], stu[l]
			l++
			less++
		} else if stu[l].grade == base {
			l++ // 相等只需要移动指针
		} else {
			more--
			stu[l], stu[more] = stu[more], stu[l]
		}
	}

	return []int{less + 1, more - 1}
}

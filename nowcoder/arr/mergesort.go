package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 归并实现 --> 归并稳定，是该题的解决方案
// HJ68 成绩排序

func main() {
	n := 0
	sort := 0
	fmt.Scan(&n)
	fmt.Scan(&sort)

	reader := bufio.NewReader(os.Stdin)
	stus := make([]student, n)
	for i := 0; i < n; i++ {
		msg, _ := reader.ReadString('\n')
		msgs := strings.Fields(msg)
		grade, _ := strconv.Atoi(msgs[1])
		stus[i] = student{
			name:  msgs[0],
			grade: grade,
		}
	}
	mergesort(stus, 0, n-1)

	for i := range stus {
		if sort == 1 {
			fmt.Println(stus[i].name, stus[i].grade)
		} else {
			fmt.Println(stus[n-1-i].name, stus[n-1-i].grade)
		}
	}

}

func mergesort(stus []student, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)/2
	mergesort(stus, l, mid)
	mergesort(stus, mid+1, r)

	merge(stus, l, mid, r)
}

func merge(stus []student, l, m, r int) {
	helper := make([]student, r-l+1)
	p1 := l
	p2 := m + 1
	i := 0 // helper的Index

	for p1 <= m && p2 <= r {
		if stus[p1].grade > stus[p2].grade {
			helper[i] = stus[p2]
			p2++
		} else {
			helper[i] = stus[p1]
			p1++
		}
		i++
	}
	for p1 <= m {
		helper[i] = stus[p1]
		i++
		p1++
	}
	for p2 <= r {
		helper[i] = stus[p2]
		i++
		p2++
	}

	for i := 0; i < len(helper); i++ {
		stus[l+i] = helper[i]
	}
}

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
// 描述
// 给定一些同学的信息（名字，成绩）序列，请你将他们的信息按照成绩从高到低或从低到高的排列,相同成绩
//
// 都按先录入排列在前的规则处理。
//
// 例示：
// jack      70
// peter     96
// Tom       70
// smith     67
// 从高到低  成绩
// peter     96
// jack      70
// Tom       70
// smith     67
// 从低到高
// smith     67
// jack      70
// Tom       70
// peter     96
// 注：0代表从高到低，1代表从低到高
//
// 数据范围：人数：1≤n≤200
// 进阶：时间复杂度：O(nlogn) ，空间复杂度：O(n)
// 输入描述：
// 第一行输入要排序的人的个数n，第二行输入一个整数表示排序的方式，之后n行分别输入他们的名字和成绩，以一个空格隔开
//
// 输出描述：
// 按照指定方式输出名字和成绩，名字和成绩之间以一个空格隔开
//
// 示例1
// 输入：
// 3
// 0
// fang 90
// yang 50
// ning 70
// 复制
// 输出：
// fang 90
// ning 70
// yang 50
// 复制
// 示例2
// 输入：
// 3
// 1
// fang 90
// yang 50
// ning 70
// 复制
// 输出：
// yang 50
// ning 70
// fang 90

type student struct {
	name  string
	grade int
}

func heapsort() {
	//func main() {
	n := 0
	sort := 0

	fmt.Scan(&n)
	fmt.Scan(&sort)

	stu := make([]student, n)
	sc := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		msg, _ := sc.ReadString('\n')
		msgs := strings.Fields(msg)
		grade, _ := strconv.Atoi(msgs[1])
		stu[i] = student{
			name:  msgs[0],
			grade: grade,
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(stu, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		stu[0], stu[i] = stu[i], stu[0]
		heapify(stu, i, 0)
	}

	for i := range stu {
		if sort == 0 { // 从大到小，逆序输出即可
			fmt.Println(stu[n-1-i].name, stu[n-1-i].grade)

		} else { // 从小打到，正序
			fmt.Println(stu[i].name, stu[i].grade)
		}
	}
}

// 堆化,但是堆排不是稳定的，注意了
func heapify(stu []student, n, i int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2
		if left < n && stu[left].grade > stu[largest].grade {
			largest = left
		}
		if right < n && stu[right].grade > stu[largest].grade {
			largest = right
		}

		if largest == i { // 说明没有比父节点还大的子节点，停止
			break
		}

		stu[i], stu[largest] = stu[largest], stu[i]
		i = largest
	}
}

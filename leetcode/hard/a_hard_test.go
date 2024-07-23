package hard

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// 41. 缺失的第一个正数
func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "missing", args: args{nums: []int{1, 2, 0}}, want: 3},
		{name: "missing", args: args{nums: []int{3, 4, -1, 1}}, want: 2},
		//{name: "missing", args: args{nums: []int{7, 8, 9, 11, 12}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositiveV1(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 135. 分发糖果
func Test_Candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "candy", args: args{ratings: []int{1, 0, 2}}, want: 5},
		{name: "candy", args: args{ratings: []int{1, 2, 2}}, want: 4},
		{name: "candy", args: args{ratings: []int{1, 3, 2, 2, 1}}, want: 7},
		{name: "candy", args: args{ratings: []int{1, 2, 87, 87, 87, 2, 1}}, want: 13},
		{name: "candy", args: args{ratings: []int{1, 3, 4, 5, 2}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 32. 最长有效括号
func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "longest", args: args{s: ")))(()"}, want: 2},
		//{name: "longest", args: args{s: ")()())"}, want: 4},
		//{name: "longest", args: args{s: "()())"}, want: 4},
		//{name: "longest", args: args{s: "((()))())"}, want: 8},
		{name: "longest", args: args{s: "((()()))())"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesStack(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 354. 俄罗斯套娃信封问题
func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxEnvelopes", args: args{envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}}, want: 3},
		{name: "maxEnvelopes", args: args{envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}}}, want: 1},
		//[2,100],[3,200],[4,300],[5,500],[5,400],[5,250],[6,370],[6,360],[7,380]
		{name: "maxEnvelopes", args: args{envelopes: [][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("dividing", func(t *testing.T) {
		fmt.Println(math.Floor(11 / 15))
		fmt.Println(math.Floor(23 / 15))
		fmt.Println(math.Floor(30 / 15))
	})
}

// 239. 滑动窗口最大值
func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "maxSlidingWindow", args: args{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3}, want: []int{3, 3, 5, 5, 6, 7}},
		{name: "maxSlidingWindow", args: args{nums: []int{1}, k: 1}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowWithEasy(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 23.合并K个升序链表
func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "mergeKList", args: args{lists: []*ListNode{&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			}}, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			}}, &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			}}}}, want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  6,
										Next: nil,
									}},
							},
						}},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 76. 最小覆盖子串
func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "minWindow", args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 170. 交易逆序对的总数
func Test_reversePairs(t *testing.T) {
	type args struct {
		record []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "reversePairs", args: args{record: []int{9, 7, 5, 4, 6}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.record); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

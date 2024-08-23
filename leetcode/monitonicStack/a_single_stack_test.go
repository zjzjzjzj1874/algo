package monitonicStack

import (
	"reflect"
	"testing"
)

// 962. 最大宽度坡
func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxwidth", args: args{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}}, want: 7},
		{name: "maxwidth", args: args{nums: []int{2, 2, 1}}, want: 1},
		{name: "maxwidth", args: args{nums: []int{6, 0, 8, 2, 1, 5}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.nums); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 037. 行星碰撞
func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "asteroid collision", args: args{asteroids: []int{5, 10, -5}}, want: []int{5, 10}},
		{name: "asteroid collision", args: args{asteroids: []int{8, -8}}, want: []int{}},
		{name: "asteroid collision", args: args{asteroids: []int{-2, 2, 1, -2}}, want: []int{-2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollisionWithLT(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 503. 下一个更大元素 II
func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nextGreatElements", args: args{nums: []int{1, 2, 1}}, want: []int{2, -1, 2}},
		{name: "nextGreatElements", args: args{nums: []int{1, 2, 3, 4, 3}}, want: []int{2, 3, 4, -1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nextGreater", args: args{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}, want: []int{-1, 3, -1}},
		{name: "nextGreater", args: args{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}, want: []int{3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementWithMap(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1019. 链表中的下一个更大节点
func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{name: "nextLargerNodes",
			args: args{head: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}}},
			want: []int{5, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 556. 下一个更大元素 III
func Test_nextGreaterElement556(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "nextGreaterElement556", args: args{n: 12}, want: 21},
		{name: "nextGreaterElement556", args: args{n: 21}, want: -1},
		{name: "nextGreaterElement556", args: args{n: 31425}, want: 31452},
		{name: "nextGreaterElement556", args: args{n: 230241}, want: 230412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement556(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement556() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 456. 132 模式
func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "find132pattern", args: args{nums: []int{1, 2, 3, 4}}, want: false},
		{name: "find132pattern", args: args{nums: []int{3, 1, 4, 2}}, want: true},
		{name: "find132pattern", args: args{nums: []int{-1, 3, 2, 0}}, want: true},

		{name: "find132pattern", args: args{nums: []int{1, 0, 1, -4, -3}}, want: false},
		{name: "find132pattern", args: args{nums: []int{3, 5, 0, 3, 4}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1475. 商品折扣后的最终价格
func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "finalPrices", args: args{prices: []int{8, 4, 6, 2, 3}}, want: []int{4, 2, 4, 2, 3}},
		{name: "finalPrices", args: args{prices: []int{10, 1, 1, 6}}, want: []int{9, 0, 1, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 036. 逆波兰表达式求值
func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "evalRPN", args: args{tokens: []string{"2", "1", "+", "3", "*"}}, want: 9},
		{name: "evalRPN", args: args{tokens: []string{"4", "13", "5", "/", "+"}}, want: 6},
		{name: "evalRPN", args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPNWithLT(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

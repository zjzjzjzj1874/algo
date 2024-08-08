package matrix

import (
	"reflect"
	"testing"
)

// 2326. 螺旋矩阵 IV
func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		// 输入：m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
		// 输出：[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
		{name: "spiralMatrix", args: args{m: 3, n: 5,
			head: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 1, Next: &ListNode{Val: 7,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 0}}}}}}}}}}}}}},
			wantAns: [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spiralMatrix() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 59. 螺旋矩阵 II
func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//[1,2,3],[8,9,4],[7,6,5]
		{name: "generateMatrix", args: args{n: 3}, want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		// [1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]
		{name: "generateMatrix", args: args{n: 4}, want: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 54. 螺旋矩阵
func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{name: "spiralOrder", args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, wantAns: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		//[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
		//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
		{name: "spiralOrder", args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, wantAns: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spiralOrder(tt.args.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spiralOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

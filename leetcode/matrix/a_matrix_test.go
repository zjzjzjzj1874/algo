package matrix

import (
	"reflect"
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spiralOrder(tt.args.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spiralOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

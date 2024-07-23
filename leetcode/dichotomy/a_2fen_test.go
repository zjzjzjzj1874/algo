package dichotomy

import (
	"reflect"
	"testing"
)

// 153. 寻找旋转排序数组中的最小值
func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "findMin", args: args{nums: []int{3, 1, 2}}, want: 1},
		//{name: "findMin", args: args{nums: []int{0, 1, 2, 3}}, want: 0},
		{name: "findMin", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		//{name: "findMin", args: args{nums: []int{11, 13, 15, 17}}, want: 11},
		//{name: "findMin", args: args{nums: []int{5, 6, 7, 1, 2, 3, 4}}, want: 1},
		//{name: "findMin", args: args{nums: []int{6, 7, 8, 1, 2, 3}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinV1(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 162. 寻找峰值
func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findPeakElement", args: args{nums: []int{1, 2, 3, 1}}, want: 2},
		{name: "findPeakElement", args: args{nums: []int{1, 2, 4, 6, 13}}, want: 4},
		{name: "findPeakElement", args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElementV1(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 2529. 正整数和负整数的最大计数
func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maximum", args: args{nums: []int{-2, -1, -1, 1, 2, 3}}, want: 3},
		{name: "maximum", args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}}, want: 3},
		{name: "maximum", args: args{nums: []int{5, 20, 66, 1314}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 33. 搜索旋转排序数组
func Test_search33(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{name: "search", args: args{nums: []int{1}, target: 0}, want: -1},
		{name: "search", args: args{nums: []int{5, 1, 2, 3, 4}, target: 1}, want: 1},
		{name: "search", args: args{nums: []int{8, 1, 2, 3, 4, 5, 6, 7}, target: 6}, want: 6},
		{name: "search", args: args{nums: []int{1, 3, 5}, target: 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search33(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search33() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 704. 二分查找
func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "search", args: args{nums: []int{5}, target: 5}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchV1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 35. 搜索插入位置
func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "searchInsert", args: args{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{name: "searchInsert", args: args{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{name: "searchInsert", args: args{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
		{name: "searchInsert", args: args{nums: []int{1}, target: 7}, want: 1},
		{name: "searchInsert", args: args{nums: []int{1}, target: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsertV1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "searchRange", args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, want: []int{3, 4}},
		{name: "searchRange", args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6}, want: []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 073. 爱吃香蕉的狒狒
func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minRatingSpeed", args: args{piles: []int{3, 6, 7, 11}, h: 8}, want: 4},
		{name: "minRatingSpeed", args: args{piles: []int{30, 11, 23, 4, 20}, h: 5}, want: 30},
		{name: "minRatingSpeed", args: args{piles: []int{30, 11, 23, 4, 20}, h: 6}, want: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

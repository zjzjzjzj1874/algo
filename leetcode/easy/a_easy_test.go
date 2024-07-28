package easy

import (
	"reflect"
	"testing"
)

// 888. 公平的糖果交换
func Test_fairCandySwap(t *testing.T) {
	type args struct {
		aliceSizes []int
		bobSizes   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "fairCandySwap", args: args{aliceSizes: []int{1, 1}, bobSizes: []int{2, 2}}, want: []int{1, 2}},
		{name: "fairCandySwap", args: args{aliceSizes: []int{1, 2}, bobSizes: []int{2, 3}}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwapWithLT(tt.args.aliceSizes, tt.args.bobSizes); !reflect.DeepEqual(got, tt.want) {
				//if got := fairCandySwapOn(tt.args.aliceSizes, tt.args.bobSizes); !reflect.DeepEqual(got, tt.want) {
				//if got := fairCandySwap(tt.args.aliceSizes, tt.args.bobSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 383. 赎金信
func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "canConstruct", args: args{ransomNote: "a", magazine: "b"}, want: false},
		{name: "canConstruct", args: args{ransomNote: "aa", magazine: "ab"}, want: false},
		{name: "canConstruct", args: args{ransomNote: "aa", magazine: "aab"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 219. 存在重复元素 II
func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "containsNearbyDuplicate", args: args{nums: []int{1, 2, 3, 1}, k: 3}, want: true},
		{name: "containsNearbyDuplicate", args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2}, want: false},
		{name: "containsNearbyDuplicate", args: args{nums: []int{99, 99}, k: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicateHash(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 67. 二进制求和
func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{name: "addBinary", args: args{a: "11", b: "1"}, want: "100"},
		//{name: "addBinary", args: args{a: "1011", b: "1010"}, want: "10101"},
		{name: "addBinary", args: args{a: "1", b: "111"}, want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 075. 数组的相对排序
func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "relativeSortArray", args: args{arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, arr2: []int{2, 1, 4, 3, 9, 6}}, want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArrayWithLT(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				//if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 389. 找不同
func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "findTheDifference", args: args{s: "abcd", t: "abcde"}, want: 'e'},
		{name: "findTheDifference", args: args{s: "", t: "y"}, want: 'y'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifferenceWithSum(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 448. 找到所有数组中消失的数字
func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "findDisappearedNumbers", args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}}, want: []int{5, 6}},
		{name: "findDisappearedNumbers", args: args{nums: []int{1, 1}}, want: []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 120. 寻找文件副本
func Test_findRepeatDocument(t *testing.T) {
	type args struct {
		documents []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findRepeat", args: args{documents: []int{2, 5, 3, 0, 5, 0}}, want: 0}, // want：0 、5
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatDocument(tt.args.documents); got != tt.want {
				t.Errorf("findRepeatDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

package easy

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// 605. 种花问题
func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "canPlaceFlowers", args: args{flowerbed: nil, n: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 001. 两数相除
func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "divide", args: args{a: 15, b: 2}, want: 7},
		{name: "divide", args: args{a: 2, b: 2}, want: 1},
		{name: "divide", args: args{a: 7, b: -3}, want: -2},
		{name: "divide", args: args{a: -2147483648, b: -1}, want: 2147483647},
		{name: "divide", args: args{a: -2147483648, b: 4}, want: -536870912},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 3. 罗马数字转整数
func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "romanToInt", args: args{s: "IX"}, want: 9},
		{name: "romanToInt", args: args{s: "LVIII"}, want: 58},
		{name: "romanToInt", args: args{s: "MCMXCIV"}, want: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 290. 单词规律
func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "wordPattern", args: args{pattern: "abba", s: "dog cat cat dog"}, want: true},
		{name: "wordPattern", args: args{pattern: "abba", s: "dog cat cat fish"}, want: false},
		{name: "wordPattern", args: args{pattern: "abba", s: "dog dog dog dog"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1. 两数之和
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twosumTwoPtr(nums, target))

	//nums = []int{3, 2, 4}
	//target = 6
	//fmt.Println("twosumHash == ", twosumHash(nums, target))

	nums = []int{3, 3}
	target = 6
	bubbleRes := twosumBubble(nums, target)
	hashRes := twosumHash(nums, target)
	sort.Ints(bubbleRes)
	sort.Ints(hashRes)
	fmt.Printf("bubbleRes = %v, hashRes = %v,equal = %v", bubbleRes, hashRes, reflect.DeepEqual(bubbleRes, hashRes))
}

func Test_twosumHash(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "twosum", args: args{arr: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twosumHash(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosumHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 28. 找出字符串中第一个匹配项的下标
func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "hello", args: args{haystack: "hello", needle: "hello"}, want: 0},
		//{name: "hello", args: args{haystack: "butsad", needle: "sad"}, want: 3},
		{name: "hello", args: args{haystack: "BBC ABCDABABCDABCDABDE", needle: "ABCDABD"}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrWithKMP(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 27. 原地移除元素
func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Printf("nums = %v,res = %d\n", nums, removeElement(nums, val))
}

// 26. 删除有序数组中的重复项
func TestRemoveDuplicate(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicates(nums))

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicates(nums))
}

// 392. 判断子序列
func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isSubsequence", args: args{s: "abc", t: "ahbgdc"}, want: true},
		{name: "isSubsequence", args: args{s: "axc", t: "ahbgdc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 125. 验证回文串
func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isPalindrome", args: args{s: "A man, a plan, a canal: Panama"}, want: true},
		{name: "isPalindrome", args: args{s: "race a car"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeWithO1(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 9. 回文数
func Test_isPalindromeWithNum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isPalindromeWith", args: args{x: 12321}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeWithNum(tt.args.x); got != tt.want {
				t.Errorf("isPalindromeWithNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 283. 移动零
func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "moveZeroes", args: args{nums: []int{0, 1, 0, 3, 12}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}

// 268. 丢失的数字
func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "missingNumber", args: args{nums: []int{3, 0, 1}}, want: 2},
		{name: "missingNumber", args: args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 121:买卖股票的最佳时机
func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "test_max_profit", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		//{name: "test_max_profit", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "test_max_profit", args: args{prices: []int{2, 7, 2, 5, 1, 8}}, want: 7},
		{name: "test_max_profit", args: args{prices: []int{2, 1, 4}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 229. 多数元素 II
func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{name: "majority", want: []int{3}, args: args{nums: []int{3, 2, 3}}},
		{name: "majority", want: []int{1, 2}, args: args{nums: []int{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementWithMor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 169: 多数元素
func Test_majoritElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "majority", want: 3, args: args{nums: []int{3, 2, 3}}},
		{name: "majority", want: 2, args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}}},
		{name: "majority", want: 5, args: args{nums: []int{6, 5, 5}}},
		{name: "majority", want: 3, args: args{nums: []int{3, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majoritElementWithMoreV1(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 14:最长公共前缀
func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "test", args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 58:最后一个单词的长度
func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "", args: args{"   fly me   to   the moon  "}, want: 4},
		//{name: "", args: args{"Hello World"}, want: 5},
		{name: "", args: args{"a"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 20. 有效的括号
func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isValid", args: args{s: ""}, want: true},
		{name: "isValid", args: args{s: "(]"}, want: false},
		{name: "isValid", args: args{s: "()[]{}"}, want: true},
		{name: "isValid", args: args{s: "()"}, want: true},
		{name: "isValid", args: args{s: "([{}])"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 205. 同构字符串
func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isomorphic", args: args{s: "egg", t: "add"}, want: true},
		{name: "isomorphic", args: args{s: "foo", t: "bar"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 242. 有效的字母异位词
func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isAnagram", args: args{s: "anagram", t: "nagaram"}, want: true},
		{name: "isAnagram", args: args{s: "rat", t: "car"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

package middle

import (
	"fmt"
	"reflect"
	"testing"
)

// LCR 014. 字符串的排列
func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "checkInclusion", args: args{s1: "ab", s2: "eidbaooo"}, want: true},
		{name: "checkInclusion", args: args{s1: "ab", s2: "eidboaoo"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 47. 前 K 个高频元素
func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "topKFrequent", args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{name: "topKFrequent", args: args{nums: []int{1}, k: 1}, want: []int{1}},
		{name: "topKFrequent", args: args{nums: []int{1, 2}, k: 2}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentWithCount(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				//if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 15. 三数之和 https://leetcode.cn/problems/3sum/description/
func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{-1, 0, 1, 0}
	//nums := []int{1, 1, -2}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

// 167. 两数之和 II - 输入有序数组
func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "twoSum", args: args{numbers: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 721. 账户合并
func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "accounts", args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"Mary", "mary@mail.com"},
				}},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			//[["Alex","Alex5@m.co","Alex4@m.co","Alex0@m.co"],
			//["Ethan","Ethan3@m.co","Ethan3@m.co","Ethan0@m.co"],
			//["Kevin","Kevin4@m.co","Kevin2@m.co","Kevin2@m.co"],
			//["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe2@m.co"],
			//["Gabe","Gabe3@m.co","Gabe4@m.co","Gabe2@m.co"]]
			name: "accounts", args: args{
				accounts: [][]string{
					{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
					{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
					{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
					{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
					{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
				}},
			//[["Alex","Alex0@m.co","Alex4@m.co","Alex5@m.co"],
			//["Ethan","Ethan0@m.co","Ethan3@m.co"],
			//["Gabe","Gabe0@m.co","Gabe2@m.co","Gabe3@m.co","Gabe4@m.co"],
			//["Kevin","Kevin2@m.co","Kevin4@m.co"]]...
			want: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				fmt.Println("want:", tt.want)
				fmt.Println("got:", got)
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 面试题 16.10. 生存人数
func Test_maxAliveYear(t *testing.T) {
	type args struct {
		birth []int
		death []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxAlive", args: args{birth: []int{1900, 1901, 1950}, death: []int{1948, 1951, 2000}}, want: 1901},
		{name: "maxAlive", args: args{
			birth: []int{1972, 1908, 1915, 1957, 1960, 1948, 1912, 1903, 1949, 1977, 1900, 1957, 1934, 1929, 1913, 1902, 1903, 1901},
			death: []int{1997, 1932, 1963, 1997, 1983, 2000, 1926, 1962, 1955, 1997, 1998, 1989, 1992, 1975, 1940, 1903, 1983, 1969}},
			want: 1960},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAliveYear(tt.args.birth, tt.args.death); got != tt.want {
				t.Errorf("maxAliveYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 2. 两数相加
func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{name: "l1", args: args{l1: &ListNode{
		//	Val: 9,
		//	Next: &ListNode{
		//		Val: 9,
		//		Next: &ListNode{
		//			Val: 9,
		//			Next: &ListNode{
		//				Val: 9,
		//				Next: &ListNode{
		//					Val: 9,
		//					Next: &ListNode{
		//						Val: 9,
		//						Next: &ListNode{
		//							Val:  9,
		//							Next: nil,
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//}, l2: &ListNode{
		//	Val: 9,
		//	Next: &ListNode{
		//		Val: 9,
		//		Next: &ListNode{
		//			Val: 9,
		//			Next: &ListNode{
		//				Val:  9,
		//				Next: nil,
		//			},
		//		},
		//	},
		//}}, want: &ListNode{}},
		{name: "l2", args: args{l1: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		}, l2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}}, want: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 438. 找到字符串中所有字母异位词
func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{s: "cbaebabacd", p: "abc"}, want: []int{0, 6}},
		{name: "", args: args{s: "abab", p: "ab"}, want: []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 215. 数组中的第K个最大元素
func Test_findKthLargestWithHeapSort(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findKthLargest", args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestWithHeapSort(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestWithHeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 面试题 17.14. 最小K个数
func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{name: "smallestK", args: args{arr: []int{1, 3, 5, 7, 2, 4, 6, 8}, k: 4}, wantAns: []int{1, 2, 3, 4}},
		{name: "smallestK", args: args{arr: []int{
			62577, -220, -8737, -22, -6, 59956, 5363, -16699, 0, -10603, 64, -24528, -4818, 96, 5747, 2638, -223, 37663, -390, 35778, -4977, -3834, -56074, 7, -76, 601, -1712, -48874, 31, 3, -9417, -33152, 775, 9396, 60947, -1919, 683, -37092, -524, -8, 1458, 80, -8, 1, 7, -355, 9, 397, -30, -21019, -565, 8762, -4, 531, -211, -23702, 3, 3399, -67, 64542, 39546, 52500, -6263, 4, -16, -1, 861, 5134, 8, 63701, 40202, 43349, -4283, -3, -22721, -6, 42754, -726, 118, 51, 539, 790, -9972, 41752, 0, 31, -23957, -714, -446, 4, -61087, 84, -140, 6, 53, -48496, 9, -15357, 402, 5541, 4, 53936, 6, 3, 37591, 7, 30, -7197, -26607, 202, 140, -4, -7410, 2031, -715, 4, -60981, 365, -23620, -41, 4, -2482, -59, 5, -911, 52, 50068, 38, 61, 664, 0, -868, 8681, -8, 8, 29, 412},
			k: 131},
			wantAns: []int{-61087, -60981, -56074, -48874, -48496, -37092, -33152, -26607, -24528, -23957, -23702, -23620, -22721, -21019, -16699, -15357, -10603, -9972, -9417, -8737, -7410, -7197, -6263, -4977, -4818, -4283, -3834, -2482, -1919, -1712, -911, -868, -726, -715, -714, -565, -524, -446, -390, -355, -223, -220, -211, -140, -76, -67, -59, -41, -30, -22, -16, -8, -8, -8, -6, -6, -4, -4, -3, -1, 0, 0, 0, 1, 3, 3, 3, 4, 4, 4, 4, 4, 5, 6, 6, 7, 7, 7, 8, 8, 9, 9, 29, 30, 31, 31, 38, 51, 52, 53, 61, 64, 80, 84, 96, 118, 140, 202, 365, 397, 402, 412, 531, 539, 601, 664, 683, 775, 790, 861, 1458, 2031, 2638, 3399, 5134, 5363, 5541, 5747, 8681, 8762, 9396, 35778, 37591, 37663, 39546, 40202, 41752, 42754, 43349, 50068, 52500}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := smallestK(tt.args.arr, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("smallestK() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 442. 数组中重复的数据
func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "findDuplicates", args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}}, want: []int{2, 3}},
		{name: "findDuplicates", args: args{nums: []int{1, 1, 2}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{name: "findMin", args: args{nums: []int{3, 1, 2}}, want: 1},
		{name: "findMin", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "findMin", args: args{nums: []int{11, 13, 15, 17}}, want: 11},
		{name: "findMin", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinV1(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 49. 字母异位词分组
func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "", args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, want: [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramsV2(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				//if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 179. 最大数
func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "largestNumber", args: args{nums: []int{3, 30, 34, 5, 9}}, want: "9534330"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 3. 无重复字符的最长子串
func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{s: "abcabcbb"}, want: 3},
		{name: "", args: args{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 128. 最长连续序列
func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{nums: []int{100, 4, 200, 1, 3, 2}}, want: 4},
		{name: "", args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 11. 盛最多水的容器
func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{name: "", args: args{height: []int{1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 122:买卖股票的最佳时机II
func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test_max_profit", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "test_max_profit", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
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

// 53. 最大子数组和
func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "nums", args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "nums", args: args{nums: []int{5, 4, -1, 7, 8}}, want: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

// 56. 合并区间
func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "foo", args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "foo", args: args{intervals: [][]int{{1, 4}, {4, 5}}}, want: [][]int{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 会议室解题
func Test_meetingHolder(t *testing.T) {
	type args struct {
		meets [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		////[[1,3],[2,6],[8,10],[15,18]] => 2
		////[[1,4],[2,5],[3,6],[4,7]] => 3
		////[[1,6],[2,5],[3,7],[4,7]] => 4
		{name: "minMeet", args: args{meets: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: 2},
		{name: "minMeet", args: args{meets: [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}}}, want: 3},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {3, 7}, {4, 7}}}, want: 4},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {3, 4}, {4, 7}}}, want: 3},
		{name: "minMeet", args: args{meets: [][]int{{1, 6}, {2, 5}, {2, 5}, {3, 4}, {4, 7}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := meetingRooms(tt.args.meets); got != tt.want {
			if got := meetingHolder(tt.args.meets); got != tt.want {
				t.Errorf("meetingHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 209. 长度最小的子数组
func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minSubArray", args: args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}}, want: 2},
		{name: "minSubArray", args: args{target: 4, nums: []int{1, 4, 4}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 713. 乘积小于 K 的子数组
func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "numSubarrayProductGreaterThan", args: args{nums: []int{10, 5, 2, 6}, k: 100}, want: 8},
		{name: "numSubarrayProductGreaterThan", args: args{nums: []int{1, 1, 1}, k: 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 86. 分隔链表
func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "partition", args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			x: 3,
		}, want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

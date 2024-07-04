

## 二分法

二分法非常经典，思想也非常简单，但是大家经常会写出死循环，或者找不出正确的答案，最主要的原因还是以为边界判断有问题。

所以大家注意三种常见的区间处理，双闭，双开和左闭右开。

```go
// 双闭区间
func search(nums []int,target) int {
    l := 0 // 左闭
    r := n-1 // 右闭
    
    for l <= r { // 这里最后退出的条件,l=r
        mid := (l+r)/2 // 这里没考虑溢出，根据题意不会溢出
        if mid == target {
			return mid
        }
		if nums[mid] > target {
			r = mid - 1
        }else{
			l = mid + 1
        }       
    }
	return -1
}

// 双开区间
func search(nums []int,target) int {
    l := -1 // 左开
    r := n // 右开
    
    for l+1 < r { // 这里最后退出的条件,l=r-1
        mid := (l+r)/2 // 这里没考虑溢出，根据题意不会溢出
        if mid == target {
			return mid
        }
		if nums[mid] > target {
			r = mid // 开区间
        }else{
			l = mid // 开区间
        }       
    }
	return -1
}

// 左闭右开区间
func search(nums []int,target) int {
    l := 0 // 左闭
    r := n // 右开
    
    for l < r { // 这里最后退出的条件,l=r
        mid := (l+r)/2 // 这里没考虑溢出，根据题意不会溢出
        if mid == target {
			return mid
        }
		if nums[mid] > target {
			r = mid 
        }else{
			l = mid + 1
        }       
    }
	return -1
}




```

* [704.二分查找](https://leetcode.cn/problems/binary-search/description/)
* [33.搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/description/)
* [34.在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/)
* [35.搜索插入位置](https://leetcode.cn/problems/search-insert-position/description/)
* [153.寻找旋转排序数组中的最小值](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/)
* [744.寻找比目标字母大的最小字母](https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/)
* [2529.正整数和负整数的最大计数](https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/)
* [81.搜索旋转排序数组 II](https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/description/)
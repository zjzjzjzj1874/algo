//
// Created by zjzjzjzj1874 on 2024/11/12.
//

#include "binary_search.h"

// 有序数组的二分查找,返回数组的下标
int search(const int(& a)[10],int start, int end, int target){
    int l = 0;
    int r = end; // 左闭右闭区间

    while (l<=r) {
        int mid = (l+r)/2;
        if (a[mid]==target)
            return mid;
        if (a[mid] > target)
            r = mid - 1;
        else
            l = mid + 1;
    }

    return -1;
}

int searchRecursion(const int(& a)[10],int start, int end, int target){
    // 定义递归终止条件
    if (a[start] > target || a[end] < target || start > end)
        return -1;

    int mid = (start+end)/2;
    if (a[mid] == target)
        return mid;
    else if (a[mid] > target) // 中间值比目标值大，搜索前半部分
        return searchRecursion(a, start, mid-1, target);
    else // 中间值比目标值小，搜索后半部分
        return searchRecursion(a,mid+1,end,target);
}

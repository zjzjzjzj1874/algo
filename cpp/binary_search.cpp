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

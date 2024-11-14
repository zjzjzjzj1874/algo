//
// Created by zjzjzjzj1874 on 2024/10/31.
//

#include "sort.h"

// 返回的是数组
int* partition(int *array, int l, int r) {
    int less = l-1;
    int more = r+1;
    int base = array[r]; // 取最后一个元素作为基准值，然后比对

    while (l < more) {
        if (array[l] < base) {
            less++;
            int temp = array[l];
            array[l] = array[less];
            array[less] = temp;
            l++;
        } else if (array[l] == base) {
            l++;
        } else {
            more--;
            int temp = array[l];
            array[l] = array[more];
            array[more] = temp;
        }
    }

    int result[] = {less+1, more-1};

    return result;
}

void quick_helper(int *array, int l, int r) {
    if (l < r) {
        int *arr = partition(array, l, r);
        quick_helper(array, l, arr[0]-1);
        quick_helper(array, arr[1]+1, r);
    }
}

// 快速排序
void quick_sort(int *array, int size) {
    quick_helper(array, 0, size-1);
}


// 选择排序
// 入参：arr：数组；size：数组大小；
// 思路：每一轮找出最小的数，直到数组有序
void select_sort(int *arr, int size) {
  for (int i = 0; i < size; i++) {
    int min = i; // 定义本轮最小的数据
    for (int j = i+1; j < size; j++) {
      if (arr[min] > arr[j]) {
        min = j;
      }
    }
    int temp = arr[i];
    arr[i] = arr[min];
    arr[min] = temp;
  }
}

// 冒泡排序
// 入参：arr待排序数组；size：数组大小
// 思路：依次搞定某个元素和后面其他元素的大小关系
void bubble_sort(int *arr, int size) {
  for (int i = 0; i < size; i++) {
    for (int j = i+1; j < size; j++) {
      if (arr[i] > arr[j]) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
      }
    }
  }
}
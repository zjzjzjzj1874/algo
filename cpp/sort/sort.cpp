//
// Created by zjzjzjzj1874 on 2024/10/31.
//

#include "sort.h"

// 选择排序
// 入参：arr：数组；size：数组大小；
// 思路：每一轮找出最小的数，直到数组有序
void select_sort(int *arr, int size) {
  int n = size/sizeof(arr[0]); // 数组长度
  for (int i = 0; i < n; i++) {
    int min = i; // 定义本轮最小的数据
    for (int j = i+1; j < n; j++) {
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
  int n = size/sizeof(arr[0]); // 求数组长度
  for (int i = 0; i < n; i++) {
    for (int j = i+1; j < n; j++) {
      if (arr[i] > arr[j]) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
      }
    }
  }
}
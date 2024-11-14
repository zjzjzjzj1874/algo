//
// Created by zjzjzjzj1874 on 2024/10/31.
//

#ifndef SORT_H
#define SORT_H

// 选择排序
void select_sort(int *array, int size);

// 冒泡排序
void bubble_sort(int *array, int size);

int* partition(int *array, int l, int r);
void quick_helper(int *array, int l, int r);
// 快速排序
void quick_sort(int *array, int size);

#endif //SORT_H

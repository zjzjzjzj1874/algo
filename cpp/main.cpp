//
// Created by zjzjzjzj1874 on 2024/10/18.
//

#include <iostream>
#include <vector>
#include "math/prime.h"
#include "math/rotate_image.h"
#include "math/check_magic_square.h"
#include "math/add_string.h"
#include "sort/sort.h"
#include "reverse.h"
#include "binary_search.h"

using namespace std;

// 函数声明:反转链表
void reverseListNode();

// 打印数组
void printArray(int (&)[10]);
/*
 * cpp算法的入口
 */
int main() {
    { // 快速排序
        int arr1[] = {2, 13, 17, 11, 23, 29, 31, 19, 3, 5, 7, 37};
        quick_sort(arr1, sizeof(arr1)/sizeof(arr1[0]));
        cout << "完成快速排序：" << endl;
        for (int num : arr1) {
            cout << num << "\t";
        }
        cout << endl;
    }


    const int (& bsa)[10] = {2, 3, 5, 7, 9, 10, 13, 15, 19, 23};
    cout << "查找的元素在：" << search(bsa,0,9,15) << endl;
    cout << "查找的元素在：" << search(bsa,0,9,10) << endl;
    cout << "查找的元素在：" << search(bsa,0,9,30) << endl;
    cout << "查找的元素在：" << searchRecursion(bsa,0,9,15) << endl;
    cout << "查找的元素在：" << searchRecursion(bsa,0,9,10) << endl;
    cout << "查找的元素在：" << searchRecursion(bsa,0,9,30) << endl;
    int result = searchRecursion(bsa,0,9,30);
    result == -1 ?
        cout << "数组中没找到30！" << endl:
        cout << "数组中找到了30，下标在：" << result << endl;


    // 反转链表
    reverseListNode();

    int n = 3; // 数组的大小

    // 动态分配二维数组
    int** arr = new int*[n];
    for (int i = 0; i < n; ++i) {
        arr[i] = new int[n];
    }

    // 填充数组
    int value = 1;
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < n; ++j) {
            arr[i][j] = value++;
        }
    }

    std::cout << "原始数组:" << std::endl;
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < n; ++j) {
            std::cout << arr[i][j] << " ";
        }
        std::cout << std::endl;
    }

    // 调用旋转函数
    rotate_image(arr, n);

    std::cout << "旋转后的数组:" << std::endl;
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < n; ++j) {
            std::cout << arr[i][j] << " ";
        }
        std::cout << std::endl;
    }



    string num1 = "123", num2 = "11";
    // string num1 = "56", num2 = "77";
    string sum = add_string_without_vector(num1, num2);
    // string sum = add_string_optimize(num1, num2);
    cout << num1 << "+" << num2 << "=" << sum << endl;

    const vector<vector<int>>& ms = {
        {vector<int>{4,9,2}},
        {vector<int>{3,5,7}},
        {vector<int>{8,1,6}},
    };
    bool b = check_magic_square(ms);
    cout << (b ?"是幻方":"不是幻方") << endl;

    vector<char> c = {'h','e','l','l','e','o','u','i'};
    reverseString2ptr(c);
    for (char t : c) {
        cout << t << "\t";
    }
    cout << endl;


    // 质数
    prime();

    // 打印0-100内的所有质数
    cout << "------打印0-100内的所有质数------\n";
    for (int i = 2; i <= 100; i++) {
        if (is_prime(i)) {
            cout << i << "\t";
        }
    }
    cout << endl;

    int arr1[] = {2, 13, 17, 11, 23, 29, 31, 19, 3, 5, 7, 37};
    select_sort(arr1, sizeof(arr1)/sizeof(arr1[0]));
    cout << "完成选择排序：" << endl;
    for (int num : arr1) {
        cout << num << "\t";
    }
    cout << endl;

    int arr2[] = {2, 13, 17, 11, 23, 29, 31, 19, 3, 5, 7, 37};
    bubble_sort(arr2, sizeof(arr2)/sizeof(arr2[0]));
    cout << "完成冒泡排序：" << endl;
    for (int num : arr2) {
        cout << num << "\t";
    }
    cout << endl;
}

void reverseListNode() {
    ListNode node5 = ListNode(5);
    ListNode node4 = ListNode(4, &node5);
    ListNode node3 = ListNode(3, &node4);
    ListNode node2 = ListNode(2, &node3);
    ListNode node1 = ListNode(1, &node2);

    ListNode* cur = &node1;
    while (cur) {
        cout << cur->val << "->";
        cur = cur->next;
    }
    cout << "NULL" << endl;

    ListNode* nd = reverseList(&node1);
    cout << "翻转之后链表：" << endl;
    while (nd) {
        cout << nd->val << "->";
        nd = nd->next;
    }
    cout << "NULL" << endl;
}




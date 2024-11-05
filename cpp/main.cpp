//
// Created by zjzjzjzj1874 on 2024/10/18.
//

#include <iostream>
#include <vector>
#include "math/prime.h"
#include "sort/sort.h"
#include "reverse.h"

using namespace std;

/*
 * cpp算法的入口
 */
int main() {
    vector<char> c = {'h','e','l','l','e','o','u','i'};
    reverseString(c);
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
    select_sort(arr1, sizeof(arr1));
    cout << "完成选择排序：" << endl;
    for (int num : arr1) {
        cout << num << "\t";
    }
    cout << endl;

    int arr2[] = {2, 13, 17, 11, 23, 29, 31, 19, 3, 5, 7, 37};
    bubble_sort(arr2, sizeof(arr2));
    cout << "完成冒泡排序：" << endl;
    for (int num : arr2) {
        cout << num << "\t";
    }
    cout << endl;
}

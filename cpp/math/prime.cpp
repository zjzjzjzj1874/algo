//
// Created by zjzjzjzj1874 on 2024/10/25.
//

#include <iostream>
#include "prime.h"

using namespace std;

void prime() {
    int i = 0;
    while (true) {
        cout << "请输入一个数字：(0为退出循环)\n";
        cin >> i;
        if (i == 0) {
            break;
        }

        is_prime(i) ? cout << i << "是质数。\n" : cout << i << "不是质数。\n";
    }
}

bool is_prime(int n) {
    if (n > 0 && n <= 3) {
        return true;
    } else if (n < 0) {
        return false;
    }

    for (int i = n/2; i >= 2; i-- ) {
        if (n % i == 0) {
            return false;
        }
    }

    return true;
}
//
// Created by zjzjzjzj1874 on 2024/10/18.
//

#include <iostream>
#include "math/prime.h"

using namespace std;

/*
 * cpp算法的入口
 */
int main() {
    // 质数
    prime();

    // 打印0-100内的所有质数
    cout << "------打印0-100内的所有质数------\n";
    for (int i = 2; i <= 100; i++) {
        if (is_prime(i)) {
            cout << i << "\t";
        }
    }
}

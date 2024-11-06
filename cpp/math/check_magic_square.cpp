//
// Created by zjzjzjzj1874 on 2024/11/5.
//
#include <vector>

#include "check_magic_square.h"

bool check_magic_square(const std::vector<std::vector<int>>& s){
    int n = s.size();
    // 校验是否是N*N的矩阵
    for (std::vector<int> ss: s ) {
        if (ss.size() != n) {
            return false;
        }
    }
    // 计算目标值，每行每列相加之和=>等差数列求和(1+n*n)*(n*n)/2/n,化简如下。
    int target = (1+n*n)*n/2;
//    // 校验每行
//    for (int i = 0; i < n; i++) {
//        int sum = 0;
//        for (int j = 0; j < n; j++) {
//            sum += s[i][j];
//        }
//        if (sum != target) {
//            return false;
//        }
//    }
//    // 校验每列
//    for (int i = 0; i < n; i++) {
//        int sum = 0;
//        for (int j = 0; j < n; j++) {
//            sum += s[j][i];
//        }
//        if (sum != target) {
//            return false;
//        }
//    }
//    // 校验对角线1
//    int s1 = 0;
//    for (int i = 0; i < n; i++) {
//        s1 += s[i][i];
//    }
//    if (s1 != target) {
//        return false;
//    }
//    // 校验对角线2
//    int s2 = 0;
//    for (int i = 0; i < n; i++) {
//        s2 += s[i][n-i-1];
//    }
//    if (s2 != target) {
//        return false;
//    }

    // 优化

    // 校验两条对角线
    int sumDiag1 = 0, sumDiag2 = 0;
    // 校验每行
    for (int i = 0; i < n; i++) {
        int sumCol = 0, sumRow = 0;
        for (int j = 0; j < n; j++) {
            sumRow += s[i][j];
            sumCol += s[j][i];
        }
        if (sumCol != target || sumRow != target) {
            return false;
        }

        sumDiag1 += s[i][i];
        sumDiag2 += s[i][n-i-1];
    }
    if (sumDiag1 != target || sumDiag2 != target ) {
        return false;
    }

    return true;
}
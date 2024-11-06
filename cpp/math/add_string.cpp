//
// Created by zjzjzjzj1874 on 2024/11/6.
//

#include <string>
#include <vector>
#include <unordered_map>

#include "add_string.h"

std::unordered_map<char, int> ciMap = {
     {'0', 0},
     {'1', 1},
     {'2', 2},
     {'3', 3},
     {'4', 4},
     {'5', 5},
     {'6', 6},
     {'7', 7},
     {'8', 8},
     {'9', 9},
};

std::unordered_map<int, char> icMap = {
    {0, '0'},
    {1, '1'},
    {2, '2'},
    {3, '3'},
    {4, '4'},
    {5, '5'},
    {6, '6'},
    {7, '7'},
    {8, '8'},
    {9, '9'},
};

std::string add_string (std::string& s1, std::string& s2){
    int n1 = s1.size(), n2 = s2.size();
    int n = std::max(n1, n2); // n表示较长的数据；

    // 定义一个vector的char
    std::vector<char> vc(n+1);

    int shift = 0; // 是否需要进位
    for (int i = 0; i < n; i++) {
        if (i < n1) {
            shift += ciMap[s1[n1-i-1]];
        }
        if (i < n2) {
            shift += ciMap[s2[n2-i-1]];
        }
        if (shift >= 10) {
            vc[n-i] = icMap[shift-10];
            shift /= 10;
        } else {
            vc[n-i] = icMap[shift];
            shift = 0;
        }
    }

    if (shift > 0) {
        vc[0] = icMap[shift];
    } else {
        vc.erase(vc.begin()); // 删除首个元素
    }

    return std::string(vc.begin(), vc.end());
}

// 优化一下
std::string add_string_optimize (std::string& s1, std::string& s2){
    int n1 = s1.size(), n2 = s2.size();
    int n = std::max(n1, n2); // n表示较长的数据；

    // 定义一个vector的char
    std::vector<char> vc(n+1);

    int carry = 0; // 是否需要进位
    for (int i = 0; i < n; i++) {
        if (i < n1) {
            carry += s1[n1-i-1]-'0';
        }
        if (i < n2) {
            carry += s2[n2-i-1]-'0';
        }
        carry >= 10 ? (vc[n-i] = carry%10+'0') : (vc[n-i] = carry+'0');
        carry /= 10;
    }

    if (carry > 0) {
        vc[0] = carry+'0';
    } else {
        vc.erase(vc.begin()); // 删除首个元素
    }

    return std::string(vc.begin(), vc.end());
}

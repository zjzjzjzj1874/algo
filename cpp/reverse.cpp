//
// Created by zjzjzjzj1874 on 2024/11/5.
//
#include <vector>

#include "reverse.h"

using namespace std;

void reverseString(vector<char>& c) {
    for (int i = 0; i < c.size()/2; ++i) {
        char t = c[i];
        c[i] = c[c.size()-1-i];
        c[c.size() - 1-i] = t;
    }
}
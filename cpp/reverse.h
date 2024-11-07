//
// Created by zjzjzjzj1874 on 2024/11/5.
//

#include <vector>

using namespace std;

#ifndef REVERSE_H
#define REVERSE_H

// 链表结构体定义
struct ListNode {
    int val; // 节点值
    ListNode* next; // 下一个指针
    // 构造函数
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// 翻转字符串
void reverseString(vector<char>& s);
void reverseString2ptr(vector<char>& s); // 使用双指针翻转

// 翻转链表
ListNode* reverseList(ListNode* head);

#endif //REVERSE_H

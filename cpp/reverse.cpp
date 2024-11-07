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

void reverseString2ptr(vector<char>& c) {
    int head = 0, tail = c.size()-1;
    while (head < tail) {
        char t = c[head];
        c[head] = c[tail];
        c[tail] = t;
        ++head;
        --tail;
    }
}

// 链表翻转
ListNode* reverseList(ListNode* head) {
    ListNode* curr = head;
    ListNode* prev = nullptr;

    while (curr) {
        ListNode* temp = curr -> next;
        curr -> next = prev;
        prev = curr;
        curr = temp;
    }

    return prev;
}
//
// Created by zjzjzjzj1874 on 2024/11/14.
//

#include <iostream>

#include "tree.h"

// 先序遍历树,根左右
void printTreePreOrder(TreeNode* root) {
    if (!root) { // 树已经为空
        return;
    }

    std::cout << root->val << std::endl;
    printTreePreOrder(root->left);
    printTreePreOrder(root->right);
}

// 中序遍历树
void printTreeInOrder(TreeNode* root) {
    if (!root) {
        return;
    }

    printTreeMiddleOrder(root->left);
    std::cout << root->val << std::endl;
    printTreeMiddleOrder(root->right);
}

// 后序遍历树
void printTreePostOrder(TreeNode* root) {
    if (!root)
        return;

    printTreePostOrder(root->left);
    printTreePostOrder(root->right);
    std::cout << root->val << std::endl;
}
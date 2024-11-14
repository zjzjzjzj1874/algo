//
// Created by zjzjzjzj1874 on 2024/11/14.
//
#pragma once
#ifndef ALGO_TREE_H
#define ALGO_TREE_H

// 树的结构体类型
struct TreeNode {
    TreeNode* left;
    TreeNode* right;
    int val;
};

// 树的遍历有三种，先序(根左右)，中序(左根右)，后序(左右根)

// 先序遍历树
void printTreePreOrder(TreeNode* root);

// 中序遍历树
void printTreeInOrder(TreeNode* root);

// 后序遍历树
void printTreePostOrder(TreeNode* root);


class tree {

};


#endif //ALGO_TREE_H

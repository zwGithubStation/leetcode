
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/same-tree/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g isSameTree.c -o isSameTree
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <limits.h>



struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool isSameTree(struct TreeNode* p, struct TreeNode* q){
	if (p == NULL && q == NULL)
		return true;
	if ((p == NULL && q != NULL) || 
		(p != NULL && q == NULL) ||
		(p != NULL && q != NULL && p->val != q->val))
		return false;
	return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}

int main()
{
	return 0;
}


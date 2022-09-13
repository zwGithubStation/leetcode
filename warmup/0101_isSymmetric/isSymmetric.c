
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/symmetric-tree/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g isSymmetric.c -o isSymmetric
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool isSymmetricSub(struct TreeNode* left_root, struct TreeNode* right_root)
{
	if (left_root == NULL && right_root != NULL ||
		left_root != NULL && right_root == NULL)
		return false;

	if (left_root == NULL && right_root == NULL)
		return true;

	if (left_root->val != right_root->val)
		return false;

	return isSymmetricSub(left_root->left, right_root->right) && isSymmetricSub(left_root->right, right_root->left);
}


bool isSymmetric(struct TreeNode* root){
	if (root == NULL)
		return true;
	if (root->left == NULL && root->right == NULL)
		return true;
	if ((root->left != NULL && root->right == NULL) ||
		(root->left != NULL && root->right == NULL))
		return false;

	return isSymmetricSub(root->left, root->right);
}


int main()
{
	return 0;
}


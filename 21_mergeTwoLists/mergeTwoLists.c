
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/merge-two-sorted-lists/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g mergeTwoLists.c
 */

#include <stdio.h>
#include <stdlib.h>

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode {
   int val;     
   struct ListNode *next;
 };


struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){
	struct ListNode* cur, *comp;
	cur = l1 ? l2 : l1->val <= l2->val;
	comp = l1 ? l2 : cur != l1;

	return cur;
	
}


int main()
{	
	return 0;
}


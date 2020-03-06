
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
	struct ListNode* result = NULL;
	struct ListNode* cur, *comp, *pre, *post;

	if (l1 == NULL)
		return l2;

	if (l2 == NULL)
		return l1;

	if (l2 == NULL && l1 == NULL)
		return NULL;
	
	cur = l1->val <= l2->val ? l1 : l2;
	comp = cur != l1 ? l1 : l2; 
	result = cur;
	while (cur != NULL)
	{
		if (comp == NULL || (comp != NULL && cur->val <= comp->val))
		{
			pre = cur;
			cur = cur->next;
			
		}
		else
		{
			post = cur; //record next loop cmp-to-who
			pre->next = comp;
			cur = comp;
			comp = post;
		}
	}
	
	return result;
	
}


int main()
{	
	return 0;
}


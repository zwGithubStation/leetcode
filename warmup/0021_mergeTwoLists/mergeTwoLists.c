
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/merge-two-sorted-lists/

   solve  : 1.self-solve(two lists, cur always towards to lower list's current head, 
   						 comp always towards to higher list's current head, 
   						 usd pre/post pointers facilitate the cur/comp change,
   						 make sure when cur is searched over, comp's tail could be included)
			2.official(iteration)

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

//self-solve
/*
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

	if (comp != NULL)
	{
		pre->next = comp;
	}
	
	return result;
	
}
*/

//official
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){
	struct ListNode* result = (struct ListNode *)malloc(sizeof(struct ListNode));
	struct ListNode* cur = result;
		
	result->next = NULL;
	while (l1 != NULL || l2 != NULL)
	{
		if (l1 == NULL)
		{
			cur->next = l2;
			return result->next;
		}

		if (l2 == NULL)
		{
			cur->next = l1;
			return result->next;
		}

		if (l1->val <= l2->val)
		{
			cur->next = l1;
			l1 = l1->next;
		}
		else
		{
			cur->next = l2;
			l2 = l2->next;
		}

		cur = cur->next;
	}

	return result->next;
}


int main()
{	
	return 0;
}



/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/swap-nodes-in-pairs/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g swapPairs.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

//Definition for singly-linked list.
struct ListNode {
  int val;
  struct ListNode *next;
};


struct ListNode* swapPairs(struct ListNode* head){
	struct ListNode* new_head = NULL;
	struct ListNode* left, *right, *cur_head, *cur_tail, *ret;

	if (head == NULL || head->next == NULL)
		return head;

	new_head = (struct ListNode* )malloc(sizeof(struct ListNode));
	new_head->next = head;
	cur_head = new_head;
	left = head;
	right = head->next;

	while (true)
	{
		cur_tail = right->next;
		cur_head->next = right;
		left->next = cur_tail;
		right->next = left;

		if (cur_tail == NULL)
		{
			ret = new_head->next;
			free(new_head);
			return ret;
		}
			
		else if (cur_tail->next == NULL)
		{
			ret = new_head->next;
			free(new_head);
			return ret;
		}
		else
		{
			cur_head = left;
			left = cur_head->next;
			right = left->next;
		}
	}

	ret = new_head->next;
	free(new_head);
	return ret;
}

int main()
{	
	return 0;
}


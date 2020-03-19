
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g removeNthFromEnd.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>


//Definition for singly-linked list.
struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
	struct ListNode* nodes[10000];
	int i,len;
	struct ListNode* ret = head;
	
	i = 0;
	while (head != NULL)
	{
		nodes[i++] = head;
		head = head->next;
	}
	len = i;

	if (n == 1)
	{
		if (ret->next == NULL)
			return NULL;
		
		nodes[len-2]->next = NULL;
		return ret;
	}

	if (n == len)
		return ret->next;

	nodes[len-n-1]->next = nodes[len-n+1];
	return ret;
}

int main()
{
	return 0;
}


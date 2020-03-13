
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/add-two-numbers/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g addTwo.c -o addTwo
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
	struct ListNode* ret, *ret_cur, *cur_pre;
	struct ListNode* tail;
	bool overflow = false;

	ret_cur = (struct ListNode*)malloc(sizeof(struct ListNode));
	ret = ret_cur;
	
	while (l1 != NULL && l2 != NULL)
	{
		ret_cur->val = (overflow == false) ? (l1->val + l2->val) % 10 : (l1->val + l2->val + 1) % 10;
		overflow = (overflow == false) ? ((l1->val + l2->val) >= 10 ? true : false) : ((l1->val + l2->val + 1) >= 10 ? true : false);
		ret_cur->next = (struct ListNode*)malloc(sizeof(struct ListNode));  //pre-alloc
		cur_pre = ret_cur;
		ret_cur = ret_cur->next;
		l1 = l1->next;
		l2 = l2->next;
	}

	if (l1 == NULL && l2 == NULL)
	{
		if (overflow == true)
		{
			ret_cur->val = 1;
			ret_cur->next = NULL;
			return ret;
		}
		else
		{
			free(cur_pre->next);
			cur_pre->next = NULL;
			return ret;
		}
	}

	tail = (l1 == NULL) ? l2 : l1;
	while (tail != NULL)
	{
		ret_cur->val = (overflow == false) ? tail->val : (tail->val + 1) % 10;
		overflow = (overflow == false) ? false : ((tail->val + 1) == 10 ? true : false);
		ret_cur->next = (struct ListNode*)malloc(sizeof(struct ListNode));  //pre-alloc
		cur_pre = ret_cur;
		ret_cur = ret_cur->next;
		tail = tail->next;
	}

	if (overflow == true)
	{
		ret_cur->val = 1;
		ret_cur->next = NULL;
		return ret;
	}

	free(cur_pre->next);
	cur_pre->next = NULL;
	return ret;
}


int main()
{
	return 0;
}


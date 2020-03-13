
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

   baseline solve	  :just for basic link operation, beware of you tactic

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g deleteDuplicates.c -o deleteDuplicates
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>



 struct ListNode {
      int val;
      struct ListNode *next;
 };


struct ListNode* deleteDuplicates(struct ListNode* head){
	struct ListNode* cur, *pre;
	
	if (!head)
		return NULL;

	cur = head;
	pre = cur;
	cur = cur->next;
	
	while(cur != NULL)
	{
		while(cur != NULL && cur->val == pre->val)
		{
			cur = cur->next;
		}
	
		if (cur == NULL)
		{
			pre->next = NULL;
			return head;
		}
		pre->next = cur;
		
		pre = cur;
		cur = pre->next;
	}
	return head;
}


int main()
{
	return 0;
}


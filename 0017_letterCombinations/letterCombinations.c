
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g letterCombinations.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

char matrix2[8][4] = {  {'a','b','c','\0'},
				  		{'d','e','f','\0'},
				  		{'g','h','i','\0'},
				  		{'j','k','l','\0'},
 				  		{'m','n','o','\0'},
 				  		{'p','q','r','s'},
 				  		{'t','u','v','\0'},
 				  		{'w','x','y','z'}};

void setArrayBasedByIndex(char *array, int *index, int length)
{
	
}

void increaseIndex(int *cur_idx, int *static_idx, int length)
{
	
}


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
char ** letterCombinations(char * digits, int* returnSize){
	int length,i,total,tmp;
	char *p;
	int *cur_idx, *static_idx;
	char **ret;

	if (!digits)
		return NULL;
	
	length = strlen(digits);
	tmp = 0;
	total = 1;

	static_idx = (int *)malloc(sizeof(int)*length);
	
	while (tmp < length)
	{
		if (digits[tmp] == '7' || digits[tmp] == '9')
		{
			total *= 4;
			static_idx[tmp] = 4;
		}
		else
		{
			total *= 3;
			static_idx[tmp] = 3;
		}

		tmp++;
	}

	cur_idx = (int *)malloc(sizeof(int)*length);
	memset(cur_idx, 0, sizeof(int)*length);

	*returnSize = total;

	ret = (char **)malloc(sizeof(char *)*total);
	for (i = 0; i < length; i++)
	{
		ret[i] = (char *)malloc(sizeof(char)*(length+1));
		memset(ret[i], 0, sizeof(char)*(length+1));
	}

	while (total > 0)
	{
		setArrayBasedByIndex(ret[total-1], cur_idx, length);
		increaseIndex(cur_idx, static_idx, length);
		total--;
	}

	free(cur_idx);
	free(static_idx);

	return ret;
}

int main()
{
	return 0;
}


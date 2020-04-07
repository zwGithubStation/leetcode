
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
#include <stdbool.h>

char matrix[8][4] = {  {'a','b','c','\0'},
				  		{'d','e','f','\0'},
				  		{'g','h','i','\0'},
				  		{'j','k','l','\0'},
 				  		{'m','n','o','\0'},
 				  		{'p','q','r','s'},
 				  		{'t','u','v','\0'},
 				  		{'w','x','y','z'}};

void setArrayBasedByIndex(char *array, int *index, int *matrix_idx, int length)
{
	int i;
	for (i = 0; i < length; i++)
	{
		printf("cur[%d] matrix_idx:%d, idx:%d\n", matrix_idx[i], index[i]);
		printf("element[%d] set to:%c\n", i, matrix[matrix_idx[i]][index[i]]);
		array[i] = matrix[matrix_idx[i]][index[i]];
	}

	return;
}

void increaseIndex(int *cur_idx, int *static_idx, int length)
{
	int i = length-1;
	while (true)
	{
		if (cur_idx[i] == static_idx[i]-1)
		{
			cur_idx[i] = 0; //!!!!!
			i--;
			continue;
		}

		cur_idx[i] = cur_idx[i]+1;
		break;
	}

	return;
}


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
char ** letterCombinations(char * digits, int* returnSize){
	int length,i,total,tmp;
	char *p;
	int *cur_idx, *static_idx, *matrix_idx;
	char **ret;

	if (!digits)
		return NULL;
	
	length = strlen(digits);
	tmp = 0;
	total = 1;

	static_idx = (int *)malloc(sizeof(int)*length);
	matrix_idx = (int *)malloc(sizeof(int)*length);

	printf("1\n");
	
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

		matrix_idx[tmp] = digits[tmp] - '2';
		printf("one index:%d\n", matrix_idx[tmp]);
		tmp++;

		printf("2\n");
	}

	cur_idx = (int *)malloc(sizeof(int)*length);
	memset(cur_idx, 0, sizeof(int)*length);

	*returnSize = total;

	printf("3 total=%d\n", total);

	ret = (char **)malloc(sizeof(char *)*total);
	for (i = 0; i < total; i++)
	{
		ret[i] = (char *)malloc(sizeof(char)*(length+1));
		memset(ret[i], 0, sizeof(char)*(length+1));
	}

	printf("4\n");

	while (total > 0)
	{
		printf("5\n");
		setArrayBasedByIndex(ret[total-1], cur_idx, matrix_idx, length);
		if (total != 1)  //!!!!!
			increaseIndex(cur_idx, static_idx, length); 
		total--;
		printf("6\n");
	}

	free(cur_idx);
	free(static_idx);

	printf("7\n");

	return ret;
}

int main()
{
	char test[] = "2345";
	char **ret;
	int ret_cnt;

	ret =  letterCombinations(test, &ret_cnt);

	printf("cnt = %d", ret_cnt);
	return 0;
}


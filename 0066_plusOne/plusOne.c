
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/plus-one/


   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g plusOne.c -o plusOne
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <stdbool.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
	int i,length;
	int *result;
	bool isMax = true;

	for (i = 0; i <digitsSize; i++)
	{
		if (digits[i] != 9)
		{
			isMax = false;
			break;
		}
	}

	length = (isMax == true) ? digitsSize+1 : digitsSize;
	result = (int *)malloc(sizeof(int)*length);
	*returnSize = length;

	if (isMax == true)
	{
		memset(result, 0, sizeof(int)*length);
		result[0] = 1;
		return result;
	}

	memcpy(result, digits, sizeof(int)*length);

	i = length-1;
	while(result[i] == 9)
	{
		result[i--] = 0;
	}

	result[i] += 1;
	return result;
}


int main()
{
	int test_array[10] = {2, 2, 3, 4, 5, 6, 7};
	int *result = NULL;
	int returnSize,i;
	result = plusOne(test_array, 7, &returnSize);
	for (i = 0; i < returnSize; i++)
	{
		printf("%d ", result[i]);
	}
	printf("\n");
	free(result);
	return 0;
}


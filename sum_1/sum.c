
/*
 * Copyright (C) ZWP
 * Problem: Sum https://leetcode-cn.com/problems/two-sum/
 */

#include <stdio.h>
#include <stdlib.h>

/* baseline */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
	int i,j,match_left,match_right;
	int *result_array = (int *)malloc(sizeof(int)*2);
	
	for (i = 0 ; i < numsSize-1; i++)
	{
		match_left = nums[i];
		match_right = target - match_left;
		for (j = i+1; j < numsSize; j++)
		{
			if (nums[j] == match_right)
			{
				result_array[0] = i;
				result_array[1] = j;
				*returnSize = 2;
				return result_array;
			}
		}
	}

	free(result_array);
	*returnSize = 0;
	return NULL;
}


int main()
{
	int test_array[10] = {2, 2, 3, 4, 5, 6, 7};
	int *result = NULL;
	int returnSize;
	result = twoSum(test_array, 7, 7, &returnSize);
	printf("%d %d\n", result[0], result[1]);
	free(result);
	return 0;
}


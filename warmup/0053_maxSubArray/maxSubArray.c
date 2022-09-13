
/*
 * Copyright (C) ZWP

 * Problem: Sum https://leetcode-cn.com/problems/two-sum/

   caution: only baseline; d&c and dp solution is in other project(Alg_pro_solve\Chpater_04\MaxConsecutiveArray)

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g maxSubArray.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

//brute: calc all possible subArray's sum
int maxSubArray(int* nums, int numsSize){
	int i,j;
	int cur_sum,sum = INT_MIN;
	for (i = 0; i < numsSize; i++)
	{
		if (sum < nums[i])
		{
			sum = nums[i];
		}
		cur_sum = nums[i];
		
		for (j = i+1; j < numsSize; j++)
		{
			cur_sum = cur_sum + nums[j];
			if (cur_sum > sum)
			{
				sum = cur_sum;
			}
		}
	}

	return sum;
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


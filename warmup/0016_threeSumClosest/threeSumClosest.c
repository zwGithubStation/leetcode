
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/3sum-closest/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g threeSumClosest.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int myCompar(const void *a, const void *b)
{
	return *((int *)a) - *((int *)b);
}

int threeSumClosest(int* nums, int numsSize, int target){		
	int i,j,left,right,temp;
	int ret,diff = INT_MAX;
	
	qsort(nums, numsSize, sizeof(int), myCompar);
	
	for (i = 0; i < numsSize; i++)
	{
		left = i+1;
		right = numsSize-1;

		if (i > 0 && nums[i] == nums[i-1])
			continue;
		
		while(left < right)
		{
			temp = nums[i] + nums[left] + nums[right];
			if (temp == target)
				return target;
			else if (temp < target)
			{
				if (target-temp < diff)
				{
					diff = target-temp;
					ret = temp;
				}
				left++;
			}
			else
			{
				if (temp-target < diff)
				{
					diff = temp-target;
					ret = temp;
				}
				right--;
			}
		}
	}

	return ret;
	
}


int main()
{
	return 0;
}


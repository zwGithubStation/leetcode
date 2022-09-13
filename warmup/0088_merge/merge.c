
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/merge-sorted-array/

   baseline solve :baseline with spaceCost(n)

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g merge.c -o merge
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n){
	int *nums3 = NULL;
	int i,j,k;

	if (n == 0)
		return;

	if (m == 0)
	{
		memcpy(nums1, nums2, sizeof(int)*n);
		return;
	}
	
	nums3 = (int *)malloc(sizeof(int) * (m+n));

	i = 0;
	j = 0;
	k = 0;
	while (1)
	{
		if (j == m)
		{
			memcpy(&nums3[i], &nums2[k], sizeof(int)*(n-k));
			break;
		}

		if (k == n)
		{
			memcpy(&nums3[i], &nums1[j], sizeof(int)*(m-j));
			break;
		}
		
		if (nums1[j] < nums2[k])
		{
			nums3[i++] = nums1[j++];
		}
		else
		{
			nums3[i++] = nums2[k++];
		}
	}

	memcpy(nums1, nums3, sizeof(int)*(m+n));
	free(nums3);
	return;
}
int main()
{
	return 0;
}


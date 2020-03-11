
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g removeDuplicates.c -o removeDuplicates
 */

#include <stdio.h>
#include <stdlib.h>

int removeDuplicates(int* nums, int numsSize){
	int i,j;
	
	i = 0;
	j = 1;
	while (i < numsSize)
	{
		while (j < numsSize && nums[j] == nums[i])
		{
			j++;
		}

		if (j == numsSize)
		{
			return i+1;
		}

		i++;
		nums[i] = nums[j];
	}
}


int main()
{
	int test_array1[20] = {1,2,2,2,3,3,4,5,6,7,7,7};
	int test_array2[20] = {1,2,2,2,3,3,4,5,6,7,7,7,8};

	int result1 = removeDuplicates(test_array1, 12);
	int result2 = removeDuplicates(test_array1, 13);
	
	printf("%d %d\n", result1, result2);

	return 0;
}


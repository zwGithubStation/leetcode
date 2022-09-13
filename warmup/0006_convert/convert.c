
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/zigzag-conversion/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g convert.c -o convert
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <limits.h>

char * convert(char * s, int numRows){
	int len,column,i,j,k;
	char *ret;
	int high_tmp,low_tmp,mid_tmp1,mid_tmp2;
	
	if (!s)
		return NULL;
	
	len = strlen(s);
	j = 0;
	ret = (char *)malloc(sizeof(char)*(len+1));

	if (numRows == 1 || len <= numRows)
	{
		memcpy(ret, s, sizeof(char)*len);
		ret[len] = '\0';
		return ret;
	}

	column = 0;
	while (column*(numRows - 1) < len)
		column += 2;

	k = 0;

	for (i = 0; i < numRows; i++)
	{
		for (j = 0; j <= column; j=j+2)
		{
			if (i == 0)
			{
				high_tmp = j*(numRows-1);
				if (high_tmp < len)
					ret[k++] = s[high_tmp];
			}
			else if (i == numRows-1)
			{
				low_tmp = (j+1)*(numRows-1);
				if (low_tmp < len)
				{
					ret[k++] = s[low_tmp];
				}
			}
			else
			{
				mid_tmp1 = j*(numRows-1)-i;
				mid_tmp2 = j*(numRows-1)+i;
				if (mid_tmp1 > 0 && mid_tmp1 < len)
				{
					ret[k++] = s[mid_tmp1];
				}

				if (mid_tmp2 > 0 && mid_tmp2 < len)
				{
					ret[k++] = s[mid_tmp2];
				}
			}
		}
	}

	ret[len] = '\0';
	return ret;
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


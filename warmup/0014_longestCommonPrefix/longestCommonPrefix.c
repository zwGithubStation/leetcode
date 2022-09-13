
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/longest-common-prefix/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g longestCommonPrefix.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

char * longestCommonPrefix(char ** strs, int strsSize){
	int i,j,tem,lenghtMin = 0;
	char *result = NULL;
	for (i = 0; i < strsSize; i++)
	{
		tem = strlen(strs[i]);
		if (tem < lenghtMin || lenghtMin== 0)
			lenghtMin = tem;
	}

	result = (char *)malloc(sizeof(char)*(lenghtMin+1));
	memset(result, '\0', sizeof(char)*(lenghtMin+1));
	
	for (i = 0; i < lenghtMin; i++)
	{
		for(j = 1; j < strsSize; j++)
		{
			if (strs[0][i] != strs[j][i])
				return result;
		}
		result[i] = strs[0][i];
	}

	return result;
}


int main()
{
	char *strs[20] = {"strings", "strategy", "studio", "sting", "stripper", "steam"};
	char *result = NULL;

	result = longestCommonPrefix(strs, 6);
	printf("LCP is:%s\n", result);
	
	return 0;
}


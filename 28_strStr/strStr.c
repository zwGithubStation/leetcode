
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/implement-strstr/

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g strStr.c -o strStr
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

bool isSameStr(char * haystack, char * needle, int lenght)
{
	int i;
	for (i = 0; i < lenght; i++)
	{
		if (haystack[i] != needle[i])
			return false;
	}
	return true;
}

int strStr(char * haystack, char * needle){
	int length1,length2,i,j;

	if (!strlen(needle))
		return 0;
	if (!strlen(haystack))
		return -1;

	length1 = strlen(haystack);
	length2 = strlen(needle);

	if (length2 > length1)
		return -1;

	for (i = 0; i <= length1-length2; i++)
	{
		if (haystack[i] == needle[0] && isSameStr(&haystack[i], needle, length2))
			return i;
	}

	return -1;
}


int main()
{
	//ÂÔ
	return 0;
}



/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/longest-palindromic-substring/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g longestPalindrome.c -o longestPalindrome
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <limits.h>

//force
//len_min is 2
bool isPalindrome(char * s, int len, int startPos)
{
	int i, j;
	bool ret;
	
	for (i = startPos, j = 1; i < startPos + len/2; i++, j++)
	{
		if (s[i] != s[startPos+len-j])
			return false;
	}

	return true;
}
char * longestPalindrome(char * s){
	int i,j,len;
	int right,left,max_len = 0;
	char *ret;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;

	for (i = 0; i < len; i++)
	{
		for (j = i+1; j < len; j++)
		{
			if (isPalindrome(s, j-i+1, i))
			{
				if (j-i+1 > max_len)
				{
					left = i;
					right = j;
					max_len = j-i+1;
				}
			}
		}
	}

	if (max_len == 0)
	{
		ret = (char *)malloc(sizeof(char)*2);
		ret[0] = s[0];
		ret[1] = '\0';
		return ret;
	}

	ret = (char *)malloc(sizeof(char)*(max_len+1));
	memcpy(ret, &s[left], max_len);
	ret[max_len] = '\0';
	return ret;
}

//LCS
char * longestPalindromeLCS(char * s){
	int maxLen,endPos,len,i,j;
	int **lenMatrix;
	char *reverse;
	char *ret;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;

	lenMatrix = (int **)malloc(sizeof(int *)*len);
	for (i = 0; i < len; i++)
	{
		lenMatrix[i] = (int *)malloc(sizeof(int)*len);
		memset(lenMatrix[i], 0, sizeof(int)*len);
	}

	reverse = (char *)malloc(sizeof(char)*(len+1));
	for (i = 0; i < len; i++)
		reverse[i] = s[len-i-1];
	reverse[len] = '\0';

	maxLen = 0;
	for(i = 0; i < len; i++)
	{
		for(j = 0; j < len; j++)
		{
			if (reverse[j] == s[i])
			{
				if (i == 0 || j == 0)
					lenMatrix[i][j] = 1;
				else
				{
					lenMatrix[i][j] = lenMatrix[i-1][j-1]+1;
				}

				if ((len-j-1)+ lenMatrix[i][j] -1 == i)
				{
					if (lenMatrix[i][j] > maxLen)
					{
						maxLen = lenMatrix[i][j];
						endPos = i;
					}
				}
			}
		}
	}

	for (i = 0; i < len; i++)
	{
		free(lenMatrix[i]);
	}
	free(reverse);

	if (maxLen == 0)
	{
		return "";
	}

	ret = (char *)malloc(sizeof(char)*(maxLen+1));
	memcpy(ret, &s[endPos-maxLen+1], maxLen);
	ret[maxLen] = '\0';

	return ret;
}

//force-enhanced
char * longestPalindromeFE(char * s){
	
}

//expand-mid
char * longestPalindromeEM(char * s){
	
}


//Manacher's Algorithm
char * longestPalindromeMA(char * s){
	
}


int main()
{
	char s[6] = "babad";
	char *p = longestPalindrome(s);
	printf("ret=%s\n",p);
	free(p);
	return 0;
}



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
	int maxLen,startPos,len;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;

	int lenMatrix[len][len] = {0};

	return lenMatrix[0];
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


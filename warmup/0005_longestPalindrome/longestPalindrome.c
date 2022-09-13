
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
	int *lenRecord;
	char *reverse;
	char *ret;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;
	
	lenRecord = (int *)malloc(sizeof(int)*len);
	memset(lenRecord, 0, sizeof(int)*len);

	reverse = (char *)malloc(sizeof(char)*(len+1));
	for (i = 0; i < len; i++)
		reverse[i] = s[len-i-1];
	reverse[len] = '\0';

	maxLen = 0;
	for(i = 0; i < len; i++)
	{
		for(j = len-1; j >=0; j--)
		{
			if (reverse[j] == s[i])
			{
				if (i == 0 || j == 0)
					lenRecord[j] = 1;
				else
				{
					lenRecord[j] = lenRecord[j-1]+1;
				}

			}else{
				lenRecord[j]= 0;
			}

			if (lenRecord[j] > maxLen)   
			{
				if ((len-j-1)+ lenRecord[j] -1 == i)
				{
					maxLen = lenRecord[j];
					endPos = i;
				}
			}
		}
	}

	free(lenRecord);
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
	int maxLen,endPos,len,i,j;
	int *lenRecord;
	char *ret;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;

	lenRecord = (int *)malloc(sizeof(int)*len);
	memset(lenRecord, 0, sizeof(int)*len);

	maxLen = 0;
	for (i = len-1; i>=0; i--)
	{
		for (j = len-1; j >=i; j--)
		{
			lenRecord[j] = (s[i] == s[j] && (j - i < 3 || lenRecord[j-1] == 1)) ? 1 : 0;
			if (lenRecord[j] == 1 && j-i+1 > maxLen)
			{
				maxLen = j-i+1;
				endPos = j;
			}
		}
	}

	ret = (char *)malloc(sizeof(char)*(maxLen+1));
	memcpy(ret, &s[endPos-maxLen+1], maxLen);
	ret[maxLen] = '\0';

	return ret;
}

//expand-mid
int expandCenterEven(char * s, int left, int right, int* startPos)
{
	int i,j,len;

	len = strlen(s);
	
	while (left >= 0 && right < len && s[left] == s[right])
	{
		left--;
		right++;
	}

	if (right - left == 1)
		return 0;

	*startPos = left+1;
	return right-left-1;
}

int expandCenterOdd(char * s, int left, int right, int* startPos)
{
	int i,j,len;

	len = strlen(s);
	
	while (left >= 0 && right < len && s[left] == s[right])
	{
		left--;
		right++;
	}

	*startPos = left+1;
	return right-left-1;
}


char * longestPalindromeEM(char * s){
	int maxLen,startPos,len,lenTmp,len1,len2,start1,start2,i,j;
	char *ret;

	if (!s)
		return NULL;

	len = strlen(s);
	if (len == 1)
		return s;

	maxLen = 0;
	for (i = 0; i < len; i++)
	{
		len1 = expandCenterOdd(s, i, i, &start1);
		len2 = expandCenterEven(s, i, i+1, &start2);
		lenTmp = len1 > len2 ? len1 : len2;
		if (lenTmp > maxLen)
		{
			maxLen = lenTmp;
			startPos = len1 > len2 ? start1 : start2;
		}
	}

	ret = (char *)malloc(sizeof(char)*(maxLen+1));
	memcpy(ret, &s[startPos], maxLen);
	ret[maxLen] = '\0';

	return ret;
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


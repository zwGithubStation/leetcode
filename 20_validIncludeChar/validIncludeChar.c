
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/valid-parentheses/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g validIncludeChar.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define BLANK_CHAR 'c'

/*
bool isAllBlank(char *s, int lenght)
{
	int i;
	for (i = 0; i < lenght; i++)
	{
		if (s[i] != BLANK_CHAR)
			return false;
	}

	return true;
}
*/

void chunkArray(char *s, int *lenght)
{
	int i,j,old_length = *lenght;
	i = 0;
	j = 0;

	while (i < old_length)
	{
		if (s[i] == BLANK_CHAR)
		{
			i = i+2;
			continue;
		}
		else
		{
			s[j++] = s[i++];
			continue;
		}
	}

	s[j] = '\0';
	*lenght = strlen(s);
}

bool changeToBlank(char *s, int lenght, int *blank_cnt)
{
	bool res = false;
	int i = 0;
	*blank_cnt = 0;
	while(i < lenght-1)
	{	
		if (s[i] == s[i+1])
		{
			s[i] = BLANK_CHAR;
			s[i+1] = BLANK_CHAR;
			res = true;
			i = i+2;
			*blank_cnt = *blank_cnt + 2;
			continue;
		}

		i = i+1;
	}
	return res;
}

bool isValid(char *s){
	int length = strlen(s);
	int blank_cnt = 0;
	if (length == 0)
		return true;

	if (length % 2 != 0)
		return false;

	while(1)
	{
		changeToBlank(s, length, &blank_cnt);

		if (blank_cnt == length)
			return true;

		if (blank_cnt == 0)
			return false;

		chunkArray(s, &length);
	}

	return true;
}



int main()
{
	char strs[60] = "{[([{[{}()]}])({([{()}])})({[{}]})]}";

	result = isValid(strs);
	if (result == true)
		printf("strs is valid");
	else
		printf("strs is invalid");
	
	return 0;
}


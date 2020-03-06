
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/valid-parentheses/

   solve  : baseline(each round trunk the adjacent MATCH 2 chars, and finally get a NULL string)
   			cleanway(stack)

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g validIncludeChar.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define BLANK_CHAR 'c'

#define MATCH_CHAR(x, y) ((x == '[' && y == ']') || (x == '(' && y == ')') || (x == '{' && y == '}'))

/*
//baseline
void chunkArray(char *s, int *lenght)
{
	int i,j,old_length = *lenght;
	i = 0;
	j = 0;

	printf("old string:%s\n", s);

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

	printf("after chunk string:%s\n", s);
}

void changeToBlank(char *s, int lenght, int *blank_cnt)
{
	int i = 0;
	*blank_cnt = 0;
	while(i < lenght-1)
	{	
		if (MATCH_CHAR(s[i],s[i+1]))
		{
			printf("found touch equal:%c %c\n", s[i], s[i+1]);
			s[i] = BLANK_CHAR;
			s[i+1] = BLANK_CHAR;
			i = i+2;
			*blank_cnt = *blank_cnt + 2;
			continue;
		}

		i = i+1;
	}
	
	return;
}

bool isValid(char *s){
	int length = strlen(s);
	int blank_cnt = 0;
	if (length == 0)
		return true;

	if (length % 2 != 0)
	{
		printf("not even length!\n");
		return false;
	}
		

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
*/

//cleanway
bool isValid(char *s){
	int length = strlen(s);
	char stack[length+1];
	int i,j;

	if (length == 0)
		return true;

	if (length % 2 != 0)
	{
		return false;
	}

	memset(stack, 0, sizeof(char)*(length+1));
	i = 0;
	j = 0;

	while(s[i] != '\0')
	{
		char left = s[i];
		stack[j++] = s[i++];
		if (MATCH_CHAR(left, s[i]))
		{
			stack[j--] = '\0';
		}	
	}

	return strlen(stack) == 0;
}

int main()
{
	char strs[60] = "{[([{[{}()]}])({([{()}])})({[{}]})]}";
	
	bool result = isValid(strs);
	if (result == true)
		printf("strs is valid\n");
	else
		printf("strs is invalid\n");
	
	return 0;
}


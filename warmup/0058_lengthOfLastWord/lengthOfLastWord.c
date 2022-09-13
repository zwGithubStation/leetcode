
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/length-of-last-word/

   nothing for caution

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g lengthOfLastWord.c -o lengthOfLastWord
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int lengthOfLastWord(char * s){
	int i,cnt=0;
	
	if (!s || strlen(s) == 0)
		return 0;

	i = strlen(s)-1;

	while (i >= 0 && s[i] == ' ')
		i--;

	if (i == 0 && s[i] == ' ')
		return 0;

	while (i >= 0 && s[i] != ' ')
	{
		cnt++;
		i--;
	}

	return cnt;
}


int main()
{
	char test_array[30] = "hello world for shit    ";
	int returnSize;
	returnSize = lengthOfLastWord(test_array);
	printf("%d\n", returnSize);
	return 0;
}


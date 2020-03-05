
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/palindrome-number/

   solve  : 1.baseline (rebuild into a array, judge if head equals to tail)
   			2.cleanway (break x into head-half integer and make the reversed-tail-half integer, judge if they are the same without consider the middle element)

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g palindrome.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <limits.h>

/* baseline */
/*
bool isPalindromeArray(int *array, int lenght)
{
	int i = 0;
	for (i = 0; i < lenght / 2; i++)
	{
		if (array[i] != array[lenght-i-1])
			return false;
	}

	return true;
}

bool isPalindrome(int x)
{
	int i,pos = 0;
	int lenght = 0;
	int array[10] = {0};

	if (x < 0)
		return false;

	if (x == 0)
		return true;
	
	while (x != 0)
	{
		pos = x % 10;
		array[lenght++] = pos;
		x = x / 10;
	}

	return isPalindromeArray(array, lenght);
}
*/

/*cleanway*/
bool isPalindrome(int x)
{
	int pos,tailReversed = 0;

	if (x < 0)
		return false;
	else if (x < 10)
		return true;
	else if (x % 10 == 0)
		return false;
	
	while (x > tailReversed)
	{
		pos = x % 10;
		tailReversed = tailReversed*10 + pos;
		x = x / 10;
	}

	return (x == tailReversed) || (x == tailReversed/10);
}


int main()
{
	int input;
	bool result;
	while (scanf("%d", &input) != EOF)
	{
		result = isPalindrome(input);
		if (true == result)
			printf("%d is a palindrome!\n", input);
		else
			printf("%d is not a palindrome!\n", input);
	}
	
	return 0;
}


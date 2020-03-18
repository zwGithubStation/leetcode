
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/container-with-most-water/

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g maxArea.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <limits.h>

int maxArea(int* height, int heightSize){

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



/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/climbing-stairs/

   baseline solve	  :baseline
   C hash-based solve :baseline-hash in C

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g climbStairs.c -o climbStairs
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int combine(int x, int y)
{
	int i;
	long long divide = 1;
	long long divide_by = 1;
	
	if (y == 0 || x == y)
		return 1;

	for (i = y; i > 0; i--)
	{
		divide *= (long long)x--;
		divide_by *= (long long)y--;
	}

	return (int)divide/divide_by;
}

int climbStairs(int n){
	int i = n;
	int j = 0;
	int ret = 0;
	
	while (i >= j)
	{
		ret += combine(i--, j++);
	}

	return ret;
	
}


int main()
{
	int test = 4;
	int returnSize;
	returnSize = climbStairs(test);
	printf("%d:%d\n", test, returnSize);
	return 0;
}


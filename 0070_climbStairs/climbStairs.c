
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
	long long i;
	long long x1 = (long long)x;
	long long y1 = (long long)y;
	long long ret = 1;
	
	if (y1 == 0 || x1 == y1)
	{
		return 1;
	}

	if (y1 >= x1/2)
	{
		y1 = x1 - y1;
	}

	for (i = 1; i <= y1; i++)
	{
		ret *= x1--;
		ret /= i;
	}

	return (int)ret;
	
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
	int test = 35;
	int returnSize;
	returnSize = climbStairs(test);
	printf("%d:%d\n", test, returnSize);
	return 0;
}


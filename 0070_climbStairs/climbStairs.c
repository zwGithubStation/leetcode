
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
	long long divide = 1;
	long long divide_by = 1;
	
	if (y1 == 0 || x1 == y1)
	{
		printf("C(%d,%d) = 1\n",x, y);
		return 1;
	}

	for (i = y1; i > 0; i--)
	{
		divide *= x1--;
		divide_by *= y1--;
	}

	printf("C(%d,%d) = (%lld / %lld) = %d\n",x, y, divide, divide_by, (int)(divide/divide_by));
	return (int)(divide/divide_by);
	
}

int climbStairs(int n){
	int i = n;
	int j = 0;
	int tem,ret = 0;
	
	while (i >= j)
	{
		printf("calc C(%d,%d)\n",i, j);
		tem = combine(i, j);
		printf("calc C(%d,%d) = %d\n",i, j, tem);
		ret += tem;
		i--;
		j++;
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


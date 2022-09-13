
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/container-with-most-water/

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g maxArea.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <limits.h>

//baseline
int maxArea(int* height, int heightSize){
	int i,j;
	int ver,herz,curSize,maxSize=0;
	for (i = 0; i < heightSize-1; i++)
	{
		for (j = i+1; j < heightSize; j++)
		{
			herz = j - i;
			ver = height[i] > height[j] ? height[j] : height[i];
			curSize = herz * ver;
			if (curSize > maxSize)
				maxSize = curSize;
		}
	}

	return maxSize;
}

//double ptr
int maxArea(int* height, int heightSize){
	int i,j;
	int ver,curSize,maxSize=0;

	i = 0;
	j = heightSize - 1;

	while (i < j)
	{
		ver = height[i] >= height[j] ? height[j] : height[i];
		curSize = (j - i) * ver;
		
		if (curSize > maxSize)
			maxSize = curSize;
		
		if (height[i] < height[j])
			i++;
		else
			j--;
	}

	return maxSize;
}


int main()
{
	return 0;
}


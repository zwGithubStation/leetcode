
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/sqrtx/

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g mySqrt.c -o mySqrt
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <limits.h>
#include <math.h>


//baseline

#define LEVEL_1(x) (((x) >= 0) && ((x) <= 1000000))
#define LEVEL_2(x) (((x) > 1000000) && ((x) <= 100000000))
#define LEVEL_3(x) (((x) > 100000000) && ((x) <= 2147395600))


int loopFind(int start, int end, int x)
{
	int ret = start;
	while(ret <= end && ret * ret <= x)
	{
		if (ret * ret == x)
		{
			return ret;
		}
		ret++;
	}
	
	return ret-1;
}

int binaryFind(int start, int end, int x)
{
	int target,left;
	while (start <= end)
	{
		target = start + (end-start)/2;
		//printf("start(%d) end(%d) target(%d)\n", start, end, target);
		left = target * target;
		if (left <= x)
		{
			start = target+1;
			//printf("start(%d) end(%d) target(%d)\n", start, end, target);
		}
		else
		{
			end = target-1;
			//printf("start(%d) end(%d) target(%d)\n", start, end, target);
		}
	}

	return end;
}

int mySqrt(int x){
	int i,ret;
	
	if (LEVEL_1(x))
	{
		ret = binaryFind(0, 1000, x);
		return ret;
	}
	else if (LEVEL_2(x))
	{
		ret = binaryFind(1001, 10000, x);
		return ret;
	}
	else if (LEVEL_3(x))
	{
		ret = binaryFind(10001, 46340, x);
		return ret;
	}
	else
		return 46340;
}


//sqrt(X) = exp(0.5*lnX)  ���е���Ϊe
//log�ĵ�����10ʱ����д��lg(��ӦC���Ե�log10()����)��������eʱ���Լ�д��ln(��ӦC���Ե�log()����)
//Ҫ��log2^A  �ɸ��ݻ��׹�ʽ��ת��Ϊlog(A)/log(2)
/*
int mySqrt(int x){
	int left,right;
	if (x < 2)
		return x;
	
	left = (int)exp(0.5*log(x));
	right = left+1;

	return (long long )right * (long long )right > (long long )x ? left ? right;
}
*/


int main()
{
	int x = 15498368;
	int ret = mySqrt(x);
	printf ("%d sqrt is %d\n",x, ret);
	return 0;
}


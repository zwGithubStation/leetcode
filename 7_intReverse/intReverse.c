
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/reverse-integer/

   solve  : if not big enough(<1000000000),just reverse; 
   			if big enough(>1000000000), make sure reversed head-5-integer && tail-5-integer match require.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define POSTIVE_MAX (2147483647)
#define NEGATIVE_MAX (-2147483648)

#define PRE_TEST_MIN (1000000000)

void fill_array(int *dec_array, int x, int *lenght)
{
	int i = 0;

	if (x > PRE_TEST_MIN)
	{
		int pre_test = x % 100000;
		int pre_array_test[5] = {0};
		int pre_length_test;
		fill_array(pre_array_test, pre_test, &pre_length_test);
		pre_test = rebuild_int(pre_array_test, pre_length_test);
		if (pre_test > 21474) //reversed head-5-integer bigger than legal
		{
			*lenght = 0;
			memset(dec_array, 0, sizeof(int)*10);
			return;
		}
		if (pre_test == 21474) //reversed head-5-integer equal MAX_INT, make sure tail-5-integer not overflow
		{
			int post_test = x / 100000;
			int post_array_test[5] = {0};
			int post_length_test;
			fill_array(post_array_test, post_test, &post_length_test);
			post_test = rebuild_int(post_array_test, post_length_test);
			if (post_test >= 83647)
			{
				*lenght = 0;
				memset(dec_array, 0, sizeof(int)*10);
				return;
			}
		}
	}
	
	while (x / 10 > 0)
	{
		dec_array[i++] = x % 10;
		x = x / 10;
	}

	dec_array[i++] = x % 10;
	*lenght = i;

	return;
}

int rebuild_int(int *dec_array, int lenght)
{
	int i,head_pos = 0;
	int result = 0;
	while (dec_array[head_pos] == 0 && head_pos < lenght)
		head_pos++;

	if (head_pos == lenght)
		return 0;

	for (i = head_pos; i < lenght; i++)
	{
		result += dec_array[i];
		if (i != lenght - 1)
		{
			result *= 10;
		}
	}

	return result;
}

int reverse(int x){
	int dec_array[10] = {0};
	int abs_x,length,negative_flag = 0;

	if (x == NEGATIVE_MAX)
		return 0;
	
	if (x < 0)
	{
		x = abs(x);
		negative_flag = 1;
	}

	fill_array(dec_array, x, &length);
	x = rebuild_int(dec_array, length);

	return negative_flag == 0 ? x : (0-x);
}

int main()
{
	int x1 = 2147447410;
	int x2 = -2147447410;
	int x3 = 1000000;
	int x4 = 29003100;
	int x5 = 2147483111;
	int x6 = 1534236469;
	int x7 = -2147483648;  //-2147447412
	int x8 = -2147483647;

	int x1_res = reverse(x1);
	int x2_res = reverse(x2);
	int x3_res = reverse(x3);
	int x4_res = reverse(x4);
	int x5_res = reverse(x5);
	int x6_res = reverse(x6);
	int x7_res = reverse(x7);
	int x8_res = reverse(x8);

	printf("%d reverse is %d\n", x1, x1_res);
	printf("%d reverse is %d\n", x2, x2_res);
	printf("%d reverse is %d\n", x3, x3_res);
	printf("%d reverse is %d\n", x4, x4_res);
	printf("%d reverse is %d\n", x5, x5_res);
	printf("%d reverse is %d\n", x6, x6_res);
	printf("%d reverse is %d\n", x7, x7_res);
	printf("%d reverse is %d\n", x8, x8_res);
	return 0;
}


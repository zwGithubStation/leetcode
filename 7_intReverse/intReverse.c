
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/reverse-integer/

 */

#include <stdio.h>
#include <stdlib.h>

#define MAX_RESIVABLE (2147447412)
#define MIN_RESIVABLE (-2147447412)

void fill_array(int *dec_array, int x, int *lenght)
{
	int i = 0;
	while (x / 10 > 0)
	{
		dec_array[i++] = x % 10;
		x = x / 10;
	}

	dec_array[i++] = x % 10;
	*lenght = i;

	for (i = 0; i < 10; i++)
		printf("%d ",dec_array[i]);
	printf("  length = %d\n", *lenght);
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
	
	if (x > MAX_RESIVABLE || x < MIN_RESIVABLE)
		return 0;

	if (x == MAX_RESIVABLE || x == MIN_RESIVABLE)
		return x;

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

	int x1_res = reverse(x1);
	int x2_res = reverse(x2);

	printf("%d reverse is %d\n", x1, x1_res);
	printf("%d reverse is %d\n", x2, x2_res);
	return 0;
}



/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/string-to-integer-atoi/

 * compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g atoi.c -o atoi
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define IS_VALID_DIGID(x) ((x) >= '0' && (x) <= '9')
#define IS_VALID_NON_ZERO_DIGID(x) ((x) > '0' && (x) <= '9')

int myAtoi(char * str){
	int i,len,start,end,ret;
	bool is_negative = false;
	
	if (!str)
		return 0;

	i = 0;
	len = strlen(str);

	while(i < len && str[i] == ' ')
		i++;

	if (i == len)
		return 0;
	else if (i == len-1)
	{
		if (!IS_VALID_DIGID(str[i]))
			return 0;
		return str[i]-'0';
	}
	else
	{
		if (str[i] == '-')
		{
			if (!IS_VALID_DIGID(str[i+1]))
				return 0;
			is_negative = true;
			i++;
		}	
		else if (str[i] == '+')
		{
			if (!IS_VALID_DIGID(str[i+1]))
				return 0;
			is_negative = false;
			i++;
		}	
		else if (!IS_VALID_DIGID(str[i]))
			return 0;
	}

	start = i; 
	while (i < len && IS_VALID_DIGID(str[i]))
		i++;
	end = i;

	while (str[start] == '0')
		start++;
	
	ret = 0;
	for (i = start; i < end; i++)
	{
		if (!is_negative)
		{
			if (ret == INT_MAX / 10 && str[i] - '7' > 0)
				return INT_MAX;
			else if (ret > INT_MAX / 10)
				return INT_MAX;

		}else
		{
			if (ret == INT_MAX / 10 && str[i] == '8')
				return INT_MIN;
			else if ((ret == INT_MAX / 10 && str[i] > '8') || (ret > INT_MAX / 10))
				return INT_MIN;
		}

		ret = ret*10 + (str[i] - '0');
	}

	if (!is_negative)
		return ret;
	else
		return 0 - ret;
}


int main()
{
	return 0;
}


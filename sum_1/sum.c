
/*
 * Copyright (C) ZWP

 * Problem: Sum https://leetcode-cn.com/problems/two-sum/

   baseline solve	  :baseline
   C hash-based solve :baseline-hash in C

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g sum.c -o sum
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <limits.h>

/* baseline */
/*
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
	int i,j,match_left,match_right;
	int *result_array = (int *)malloc(sizeof(int)*2);
	
	for (i = 0 ; i < numsSize-1; i++)
	{
		match_left = nums[i];
		match_right = target - match_left;
		for (j = i+1; j < numsSize; j++)
		{
			if (nums[j] == match_right)
			{
				result_array[0] = i;
				result_array[1] = j;
				*returnSize = 2;
				return result_array;
			}
		}
	}

	free(result_array);
	*returnSize = 0;
	return NULL;
}
*/

/* baseline-hash in C */
typedef enum return_code{
	OK = 0,
	FAIL = 1
}RETURN_CODE;

//could be replaced by INT_MAX
typedef enum element_status{
	FREE = 0,
	OCCUPIED = 1
}STATUS;

typedef struct element{
	STATUS 	status; 
	int 	data;
	int 	index;
}ELEMENT;

typedef struct hash_table{
	ELEMENT *table;
	int 	 size;
}HASH_TABLE;

HASH_TABLE* hask_table_init(int size)
{
	HASH_TABLE *table = NULL;

	if (size <= 0)
	{
		return NULL;
	}

	table = (HASH_TABLE *)malloc(sizeof(HASH_TABLE));
	if (NULL == table)
	{
		return NULL;
	}

	table->size = size;
	table->table = (ELEMENT *)malloc(sizeof(ELEMENT) * size);
	if (NULL == table->table)
	{
		free(table);
		return NULL;
	}

	memset(table->table, 0, sizeof(ELEMENT) * size); //default status:free(0)
	return table;
}

void hask_table_free(HASH_TABLE* table)
{
	assert(table != NULL);

	free(table->table);
	free(table);
}

int find_element(HASH_TABLE* table, int data, int* index)
{
	assert(table != NULL);

	//abs: make sure table->table[pos_suppose] not beyond valid size in negative division scenario£¬which will cause heapMem overflow
	int pos_suppose = (abs(data) == 0) ? 0 : abs(data) % table->size;  
	int search_head = pos_suppose;

	do
	{
		if (table->table[pos_suppose].status == FREE)
		{
			return FAIL;
		}

		if  (table->table[pos_suppose].data == data)
		{
			*index = table->table[pos_suppose].index;
			return OK;
		}

		pos_suppose = (pos_suppose+1) % table->size;
	}while(pos_suppose != search_head);

	return FAIL;
}

int insert_element(HASH_TABLE* table, int data, int index)
{
	assert(table != NULL);

	//abs: make sure table->table[pos_suppose] not beyond valid size in negative division scenario£¬which will cause heapMem overflow
	int pos_suppose = (abs(data) == 0) ? 0 : abs(data) % table->size;
	int search_head = pos_suppose;

	do
	{
		if (table->table[pos_suppose].status == FREE)
		{
			table->table[pos_suppose].data  = data; 
			table->table[pos_suppose].index  = index;
			table->table[pos_suppose].status = OCCUPIED;
			return OK;
		}

		pos_suppose = (pos_suppose+1) % table->size;
	}while(pos_suppose != search_head);

	return FAIL;
}


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
	int i,add_right,match_left,match_right;
	int *result_array = (int *)malloc(sizeof(int)*2);
	
	HASH_TABLE *table = hask_table_init(numsSize * 2);
	if (NULL == table)
	{
		free(result_array);
		*returnSize = 0;
		return NULL;
	}

	for (i = 0 ; i < numsSize; i++)
	{
		add_right = target - nums[i];
		if (OK == find_element(table, add_right, &match_right))
		{
			result_array[0] = match_right;
			result_array[1] = i;
			hask_table_free(table);
			*returnSize = 2;
			return result_array;
		}

		(void)insert_element(table, nums[i], i);
	}
		
	free(result_array);
	hask_table_free(table);
	*returnSize = 0;
	return NULL;
}




int main()
{
	int test_array[10] = {2, 2, 3, 4, 5, 6, 7};
	int *result = NULL;
	int returnSize;
	result = twoSum(test_array, 7, 7, &returnSize);
	printf("%d %d\n", result[0], result[1]);
	free(result);
	return 0;
}


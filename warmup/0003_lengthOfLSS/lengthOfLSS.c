
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

   baseline solve	  :baseline

   compile: gcc -fsanitize=address -fno-omit-frame-pointer -O1 -g lengthOfLSS.c -o lengthOfLSS
 */

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include <limits.h>

/* baseline-hash in C */

//could be replaced by INT_MAX
typedef enum element_status{
	FREE = 0,
	OCCUPIED = 1
}STATUS;

typedef struct element{
	STATUS 	status; 
	char 	data;
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

bool find_element(HASH_TABLE* table, char data, int *index)
{
	assert(table != NULL);

	//abs: make sure table->table[pos_suppose] not beyond valid size in negative division scenario£¬which will cause heapMem overflow
	int pos_suppose = data - '\0';  
	int search_head = pos_suppose;

	if (table->table[pos_suppose].status == FREE)
		return false;
	else
	{
		*index = table->table[pos_suppose].index;
		return true;
	}
}

void update_element(HASH_TABLE* table, int data, int index)
{
	assert(table != NULL);

	//abs: make sure table->table[pos_suppose] not beyond valid size in negative division scenario£¬which will cause heapMem overflow
	int pos_suppose = data - '\0';  
	int search_head = pos_suppose;

	table->table[pos_suppose].data  = data; 
	table->table[pos_suppose].index  = index;
	table->table[pos_suppose].status = OCCUPIED;
	return;	
}

void delete_element(HASH_TABLE* table, int data)
{
	assert(table != NULL);

	int pos_suppose = data - '\0';  
	assert(table->table[pos_suppose].status == OCCUPIED);

	table->table[pos_suppose].status = FREE;
	return;	
}

#define MAX(x, y) ((x) > (y) ? (x) : (y))

int lengthOfLongestSubstring(char * s){
	HASH_TABLE *table = hask_table_init(256);
	int i,j,ans,index,length;

	i = 0;
	j = 0;
	ans = 0;
	length = strlen(s);

	if (length == 0)
		return 0;

	while (i < length && j < length)
	{
		if (!find_element(table, s[j], &index))
		{
			update_element(table,s[j],j+1);
			ans = MAX(ans, j-i+1);
		}
		else if (index < i)
		{
			update_element(table,s[j],j+1);
			ans = MAX(ans, j-i+1);
		}
		else
		{
			i = index;
			update_element(table, s[j],j+1);
		}
		j++;
	}

	return ans;
}

int main()
{
	return 0;
}


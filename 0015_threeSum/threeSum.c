
/*
 * Copyright (C) ZWP

 * Problem: https://leetcode-cn.com/problems/3sum/

   solve  : 

   compile: gcc -std=c99 -fsanitize=address -fno-omit-frame-pointer -O1 -g threeSum.c
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

//quickSort
void sort(int *a, int left, int right)
{
    if(left >= right)/*如果左边索引大于或者等于右边的索引就代表已经整理完成一个组了*/
    {
        return ;
    }
    int i = left;
    int j = right;
    int key = a[left];
     
    while(i < j)                               /*控制在当组内寻找一遍*/
    {
        while(i < j && key <= a[j])
        /*而寻找结束的条件就是，1，找到一个小于或者大于key的数（大于或小于取决于你想升
        序还是降序）2，没有符合条件1的，并且i与j的大小没有反转*/ 
        {
            j--;/*向前寻找*/
        }
         
        a[i] = a[j];
        /*找到一个这样的数后就把它赋给前面的被拿走的i的值（如果第一次循环且key是
        a[left]，那么就是给key）*/
         
        while(i < j && key >= a[i])
        /*这是i在当组内向前寻找，同上，不过注意与key的大小关系停止循环和上面相反，
        因为排序思想是把数往两边扔，所以左右两边的数大小与key的关系相反*/
        {
            i++;
        }
         
        a[j] = a[i];
    }
     
    a[i] = key;/*当在当组内找完一遍以后就把中间数key回归*/
    sort(a, left, i - 1);/*最后用同样的方式对分出来的左边的小组进行同上的做法*/
    sort(a, i + 1, right);/*用同样的方式对分出来的右边的小组进行同上的做法*/
                       /*当然最后可能会出现很多分左右，直到每一组的i = j 为止*/
}

void quickSort(int *a, int left, int right)
{
	int i,j,key;

	if (left >= right)
		return;
	
	i = left;
	j = right;
	key = a[left];

	while (i < j)
	{
		while (i < j && key <= a[j])
			j--;

		a[i] = a[j];

		while (i < j && key >= a[i])
			i++;

		a[j] = a[i];
	}

	a[i] = key;
	quickSort(a, left, i-1);
	quickSort(a, i+1, right);
}


/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes){
	return NULL;
}


int main()
{	
	int test[100] = {0};
	int i;

	for (i = 0; i < 100; i++)
		test[i] = rand()%100;

	for (i = 0; i < 100; i++)
		printf("%d ", test[i]);

	printf("\n");

	quickSort(test, 0, 99);

	for (i = 0; i < 100; i++)
		printf("%d ", test[i]);

	printf("\n");
	return 0;
}


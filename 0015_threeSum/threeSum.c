
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
	int retNums,i,left,right,cnt = 0;
	int **ret = NULL;

	if (numsSize < 3)
		return NULL;

	quickSort(nums, 0, numsSize-1);

	if (nums[0] > 0)
		return NULL;

	retNums = numsSize/2;

	ret = (int **)malloc(sizeof(int *)*retNums);
	returnColumnSizes = (int **)malloc(sizeof(int *)*retNums);

	for (i = 0; i < retNums; i++)
	{
		ret[i] = (int *)malloc(sizeof(int)*3);
		returnColumnSizes[i] = (int *)malloc(sizeof(int)*1);
		*(returnColumnSizes[i]) = 3;
	}

	*returnSize = 0;

	for (i = 0; i <= numsSize-3; i++)
	{
		if (nums[i] > 0)
			return ret;

		if (i > 0 && nums[i] == nums[i-1])
			continue;
			
		left = i+1;
		right = numsSize-1;
		while (left < right)
		{
			if (nums[i] + nums[left] + nums[right] == 0)
			{
				ret[cnt][0] = nums[i];
				ret[cnt][1] = nums[left];
				ret[cnt][2] = nums[right];
				cnt++;

				while (left < numsSize-1 && nums[left] == nums[left+1])
					left++;
				while (right > 0 && nums[right] == nums[right-1])
					right--;
			}else if (nums[i] + nums[left] + nums[right] < 0)
			{
				left++;
			}else
			{
				right--;
			}	
		}
	}

	if (cnt == 0)
		return NULL;

	*returnSize = cnt*3;
	return ret;
}


int main()
{	
	int test[6] = {-1,0,1,2,-1,-4};
	int i;
	int **ret;
	int **returnColumnSizes;
	int cnt;

	//for (i = 0; i < 100; i++)
	//	test[i] = rand()%100;

	//for (i = 0; i < 100; i++)
	//	printf("%d ", test[i]);

	//printf("\n");

	//quickSort(test, 0, 99);

	//for (i = 0; i < 100; i++)
	//	printf("%d ", test[i]);

	//printf("\n");

	ret = threeSum(test, 6, &cnt, returnColumnSizes);

	printf("%d\n", cnt);
	
	return 0;
}


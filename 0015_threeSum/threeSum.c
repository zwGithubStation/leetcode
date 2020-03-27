
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
    if(left >= right)/*�������������ڻ��ߵ����ұߵ������ʹ����Ѿ��������һ������*/
    {
        return ;
    }
    int i = left;
    int j = right;
    int key = a[left];
     
    while(i < j)                               /*�����ڵ�����Ѱ��һ��*/
    {
        while(i < j && key <= a[j])
        /*��Ѱ�ҽ������������ǣ�1���ҵ�һ��С�ڻ��ߴ���key���������ڻ�С��ȡ����������
        ���ǽ���2��û�з�������1�ģ�����i��j�Ĵ�Сû�з�ת*/ 
        {
            j--;/*��ǰѰ��*/
        }
         
        a[i] = a[j];
        /*�ҵ�һ������������Ͱ�������ǰ��ı����ߵ�i��ֵ�������һ��ѭ����key��
        a[left]����ô���Ǹ�key��*/
         
        while(i < j && key >= a[i])
        /*����i�ڵ�������ǰѰ�ң�ͬ�ϣ�����ע����key�Ĵ�С��ϵֹͣѭ���������෴��
        ��Ϊ����˼���ǰ����������ӣ������������ߵ�����С��key�Ĺ�ϵ�෴*/
        {
            i++;
        }
         
        a[j] = a[i];
    }
     
    a[i] = key;/*���ڵ���������һ���Ժ�Ͱ��м���key�ع�*/
    sort(a, left, i - 1);/*�����ͬ���ķ�ʽ�Էֳ�������ߵ�С�����ͬ�ϵ�����*/
    sort(a, i + 1, right);/*��ͬ���ķ�ʽ�Էֳ������ұߵ�С�����ͬ�ϵ�����*/
                       /*��Ȼ�����ܻ���ֺܶ�����ң�ֱ��ÿһ���i = j Ϊֹ*/
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


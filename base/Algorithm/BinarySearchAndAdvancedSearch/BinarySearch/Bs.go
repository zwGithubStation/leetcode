package main

//https://www.bilibili.com/video/BV1AP41137w7/?vd_source=b4d3a83d3a235cebe29e27ec23609d5e

//问题：返回有序数组中第一个>=8的数的位置；如果所有数都<8，返回数组长度

// lowerBound 返回最小的满足 nums[i] >= target 的 i
// 如果数组为空，或者所有数都 < target，则返回 nums.length
// 要求 nums 是非递减的，即 nums[i] <= nums[i + 1]

// 闭区间写法
// *********循环不变量：***********
//[left, right]区间表征当前还没有确定与target关系的区间，亦即：
// nums[left-1] < target   //left前面的元素都已经确定比target小
// nums[right+1] >= target    //right后面的元素都已经确定不比target小
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 //注意写法
		if nums[mid] < target {      //也即 left ---- mid 这段区间的元素都已经确定小于target，所以下一轮的区间缩减至[mid+1, right]
			left = mid + 1 //
		} else { //也即 mid ---- right 这段区间的元素都已经确定大于等于target， 所以下一轮的区间缩减至[left, mid+1]
			right = mid - 1 //注意不变量的语义
		}
	}
	//根据循环不变量，left前面的元素都确定比target小，最终返回left
	//又最终跳出的条件是left>right[left = right + 1]
	//所以也可以返回right+1
	return left // OR right + 1
}

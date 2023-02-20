package main

//https://www.bilibili.com/video/BV1AP41137w7/?vd_source=b4d3a83d3a235cebe29e27ec23609d5e

//涉及范围：闭区间[]，左闭右开区间[ )，开区间( )
//涉及判断条件：>=   >   <=   <

//问题：返回有序数组中第一个>= target 的数的位置；如果所有数都<target，返回数组长度

// lowerBound 返回最小的满足 nums[i] >= target 的 i
// 如果数组为空，或者所有数都 < target，则返回 nums.length
// 要求 nums 是非递减的，即 nums[i] <= nums[i + 1]

//源码 ：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/er-fen-cha-zhao-zong-shi-xie-bu-dui-yi-g-t9l9/
//1. 闭区间 [ ]
// *********循环不变量：***********
//[left, right]区间表征当前还没有确定与target关系的区间，亦即：
// nums[left-1] < target   //left前面的元素都已经确定比target小  ==========> 最终返回left
// nums[right+1] >= target    //right后面的元素都已经确定不比target==========> 最终返回right + 1
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
	//right后面的元素都已经确定不比target==========> 最终返回right + 1
	return left // OR right + 1
}

//2. 半开半闭区间 [  )
// *********循环不变量：***********
//[left, right)区间表征当前还没有确定与target关系的区间，亦即：
// nums[left-1] < target   //left前面的元素都已经确定比target小  ==========> 最终返回left
// nums[right] >= target    //right以及其后面的元素都已经确定不比target==========> 最终返回right
func lowerBound2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right { //当 left == right时， [left, right)区间为空 跳出
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left //OR right
}

//3. 开区间 (  )
// *********循环不变量：***********
//(left, right)区间表征当前还没有确定与target关系的区间，亦即：
// nums[left] < target   //left以及其前面的元素都已经确定比target小  ==========> 最终返回left+1
// nums[right] >= target    //right以及其后面的元素都已经确定不比target==========> 最终返回right
func lowerBound3(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right { // 当left + 1 = right时， (left, right)区间为空 跳出
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return left + 1 //OR right
}

//4. > 转换为 >=
// > target的第一个数   等价于   >=(target+1)的第一个数， 也即lowerBound(target+1)

//5. < 转换为 >=
// < target的第一个数(升序序列中从后往前遍历的第一个) 等价于 >=target的第一个数的前面一个数， 也即lowerBound(target) - 1

//6. <= 转换为 >=
// <=target的第一个数(升序序列中从后往前遍历的第一个) 等价于 >target的第一个数前面的一个数， 也即lowerBound(target+1) - 1

//例题

//34. https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target) //调用lowerBound or 调用lowerBound1 or 调用lowerBound2 都可以！
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1 //调用lowerBound or 调用lowerBound1 or 调用lowerBound2 都可以！
	return []int{start, end}
}

//标准库中的sort.Search  https://pkg.go.dev/sort#Search
/*
func Search(n int, f func(int) bool) int

Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true,
assuming that on the range [0, n), f(i) == true implies f(i+1) == true.

That is, Search requires that f is false for some (possibly empty) prefix of the input range [0, n) and then true for the (possibly empty) remainder;

Search returns the first true index.

If there is no such index, Search returns n.

(Note that the "not found" return value is not -1 as in, for instance, strings.Index.) Search calls f(i) only for i in the range [0, n).

A common use of Search is to find the index i for a value x in a sorted, indexable data structure such as an array or slice.
In this case, the argument f, typically a closure, captures the value to be searched for, and how the data structure is indexed and ordered.

For instance, given a slice data sorted in ascending order(升序), the call Search(len(data), func(i int) bool { return data[i] >= 23 }) returns the smallest index i such that data[i] >= 23.
If the caller wants to find whether 23 is in the slice, it must test data[i] == 23 separately.

Searching data sorted in descending order would use the <= operator instead of the >= operator.
*/

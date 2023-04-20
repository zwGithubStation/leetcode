package main

//https://www.bilibili.com/video/BV1AP41137w7/?vd_source=b4d3a83d3a235cebe29e27ec23609d5e

//涉及范围：闭区间[]，左闭右开区间[ )，开区间( )
//涉及判断条件：>=   >   <=   <

//问题：返回有序(非降序)数组中第一个>= target 的数的位置；如果所有数都<target，返回数组长度

// lowerBound 返回最小的满足 nums[i] >= target 的 i
// 如果数组为空，或者所有数都 < target，则返回 nums.length
// 要求 nums 是非递减的，即 nums[i] <= nums[i + 1]

// 源码 ：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/er-fen-cha-zhao-zong-shi-xie-bu-dui-yi-g-t9l9/
// 1. 闭区间 [ ]
// *********循环不变量：***********
// [left, right]区间表征当前还没有确定与target关系的区间，亦即：
// nums[left-1] < target   //left前面的元素都已经确定比target小  ==========> 最终返回left
// nums[right+1] >= target    //right后面的元素都已经确定不比target==========> 最终返回right + 1
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 //注意写法
		if nums[mid] < target {      //也即 left ---- mid 这段区间的元素都已经确定小于target，所以下一轮的区间缩减至[mid+1, right]
			left = mid + 1 //
		} else { //也即 mid ---- right 这段区间的元素都已经确定大于等于target， 所以下一轮的区间缩减至[left, mid-1]
			right = mid - 1 //注意不变量的语义
		}
	}
	//根据循环不变量，left前面的元素都确定比target小，最终返回left
	//right后面的元素都已经确定不比target==========> 最终返回right + 1
	return left // OR right + 1
}

// 2. 半开半闭区间 [  )
// *********循环不变量：***********
// [left, right)区间表征当前还没有确定与target关系的区间，亦即：
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

// 3. 开区间 (  )
// *********循环不变量：***********
// (left, right)区间表征当前还没有确定与target关系的区间，亦即：
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

//4. >= 转换为 >
// > target的第一个数   等价于   >=(target+1)的第一个数， 也即lowerBound(target+1)

//5. >= 转换为 <
// < target的第一个数(升序序列中从后往前遍历的第一个) 等价于 >=target的第一个数的前面一个数， 也即lowerBound(target) - 1

//6. >= 转换为 <=
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

/*
162. 寻找峰值
https://leetcode.cn/problems/find-peak-element/

峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
示例 1：
输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
提示：
1 <= nums.length <= 1000
-2e31 <= nums[i] <= 2e31 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]
*/

/* 核心问题：如果s[m] > s[m+1] 则 m 左侧必然存在峰值； 如果s[m] < s[m+1] 则 m 右侧必然存在峰值， m的取值范围：[0, len-2] */
/* 仍使用三种方案解决问题 */
/* [left, right] */
func findPeakElement(nums []int) int {
	if len(nums) == 1 { //无法在len = 1时使用左右都闭的写法，因为可能会越界
		return 0
	}
	left, right := 0, len(nums)-2 // [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/* [left, right) */
func findPeakElement2(nums []int) int {
	left, right := 0, len(nums)-1 //
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

/* (left, right) */
func findPeakElement3(nums []int) int {
	left, right := -1, len(nums)-1 // (left, right)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return left + 1
}

/*
153. 寻找旋转排序数组中的最小值
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
示例 2：
输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
示例 3：
输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。
提示：

n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
nums 中的所有整数 互不相同
nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
*/

func findMin(nums []int) int {
	return 0
}

/*
33. 搜索旋转排序数组
https://leetcode.cn/problems/search-in-rotated-sorted-array/

整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：
输入：nums = [1], target = 0
输出：-1
提示：
1 <= nums.length <= 5000
-10e4 <= nums[i] <= 10e4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10e4 <= target <= 10e4
*/

func search(nums []int, target int) int {
	return 0
}

/*
154. 寻找旋转排序数组中的最小值 II
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须尽可能减少整个过程的操作步骤。
示例 1：
输入：nums = [1,3,5]
输出：1
示例 2：
输入：nums = [2,2,2,0,1]
输出：0
提示：
n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
*/

func findMin2(nums []int) int {
	return 0
}

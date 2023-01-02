package code

import (
	"sort"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//给你一个整数 num ，返回 num 中能整除 num 的数位的数目。
//如果满足nums % val == 0 ，则认为整数 val 可以整除 nums 。
func countDigits(num int) int {
	if num < 10 {
		return 1
	}

	var ret int
	left := num / 10
	cur := num % 10
	for left > 0 {
		if cur != 0 && num%cur == 0 {
			ret++
		}
		if left < 10 {
			if left != 0 && num%left == 0 {
				ret++
			}
			break
		} else {
			cur = left % 10
			left = left / 10
		}
	}
	return ret
}

//给你一个正整数数组 nums ，对 nums 所有元素求积之后，找出并返回乘积中 不同质因数 的数目。
//注意：
//质数 是指大于 1 且仅能被 1 及自身整除的数字。
//如果 val2 / val1 是一个整数，则整数 val1 是另一个整数 val2 的一个因数。
func distinctPrimeFactors(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				set[i] = struct{}{}
				for x /= i; x%i == 0; x /= i {
				}
			}
		}
		if x > 1 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}

/*
给你一个字符串s，它每一位都是1到9之间的数字组成，同时给你一个整数k。

如果一个字符串 s的分割满足以下条件，我们称它是一个 好分割：

s中每个数位 恰好属于一个子字符串。
每个子字符串的值都小于等于k。
请你返回 s所有的 好分割中，子字符串的最少数目。如果不存在 s的好分割，返回-1。

注意：
一个字符串的 值是这个字符串对应的整数。比方说，"123"的值为123，"1"的值是1。
子字符串是字符串中一段连续的字符序列。
*/
func minimumPartition(s string, k int) int {
	var ret, curValue int
	for _, c := range s {
		curFigure := int(c - '0')
		if curFigure > k {
			return -1
		}
		curValue += curValue*10 + curFigure
		if curValue > k {
			ret++
			curValue = curFigure
		}
	}
	return ret + 1
}

/*
   给你两个正整数left 和right，请你找到两个整数num1 和num2，它们满足：

   left <= nums1 < nums2 <= right。
   nums1 和nums2都是 质数。
   nums2 - nums1是满足上述条件的质数对中的 最小值。
   请你返回正整数数组ans = [nums1, nums2]。如果有多个整数对满足上述条件，请你返回nums1最小的质数对。如果不存在符合题意的质数对，请你返回[-1, -1]。

   如果一个整数大于1，且只能被1 和它自己整除，那么它是一个质数。
*/
const mx = 1e6 + 1

var primes = make([]int, 0, 75000)

//init primes
func init() {
	notPrime := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !notPrime[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j = j + i {
				notPrime[j] = true
			}
		}
	}
	primes = append(primes, mx, mx) //防止越界
}
func closestPrimes(left int, right int) []int {
	p, q := -1, -1

	//SearchInts searches for x in a sorted slice of ints and returns the index as specified by Search.
	//The return value is the index to insert x if x is not present (it could be len(a)).
	//The slice must be sorted in ascending order.

	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}

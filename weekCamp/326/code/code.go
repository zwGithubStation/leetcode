package code

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

func minimumPartition(s string, k int) int {
	return 1
}

func closestPrimes(left int, right int) []int {
	return []int{1, 2, 3}
}

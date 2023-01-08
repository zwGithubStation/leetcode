package code

//给你一个下标从 0 开始的 环形 字符串数组 words 和一个字符串 target 。环形数组 意味着数组首尾相连。
//
//形式上， words[i] 的下一个元素是 words[(i + 1) % n] ，而 words[i] 的前一个元素是 words[(i - 1 + n) % n] ，其中 n 是 words 的长度。
//从 startIndex 开始，你一次可以用 1 步移动到下一个或者前一个单词。
//返回到达目标字符串 target 所需的最短距离。如果 words 中不存在字符串 target ，返回 -1 。
func closetTarget(words []string, target string, startIndex int) int {
	var length int = len(words)
	if words[startIndex] == target {
		return 0
	}
	for i := 1; i <= length/2; i++ {
		if words[(startIndex+i)%length] == target {
			return i
		}
		if words[(startIndex-i+length)%length] == target {
			return i
		}
	}
	return -1
}
func takeCharacters(s string, k int) int {
	return 1
}

func maximumTastiness(price []int, k int) int {
	return 1
}

func countPartitions(nums []int, k int) int {
	return 1
}

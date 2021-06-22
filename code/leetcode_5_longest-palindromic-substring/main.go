package main

import "fmt"

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
// 示例 3：
// 输入：s = "a"
// 输出："a"
// 示例 4：
// 输入：s = "ac"
// 输出："a"
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
func longestPalindrome(s string) string {
	beginRes := 0
	endRes := 0
	for i := 0; i <len(s); i++ {
		l1, r1 :=getPalindrome(s,i, i)
		l2, r2 := getPalindrome(s,i, i+1)
		maxFlagr := r1
		maxFlagl := l1
		if r1-l1 < r2-l2 {
			maxFlagr = r2
			maxFlagl = l2
		}
		if maxFlagr - maxFlagl > endRes - beginRes {
			beginRes = maxFlagl
			endRes = maxFlagr
		}
	}
	
	return s[beginRes:endRes+1]
}

func getPalindrome(s string, index1 int, index2 int) (left int, right int) {
	left = index1
	right = index2
	for{
		if left>=0 && right<len(s) && s[left] == s[right] {
			left = left - 1
			right = right + 1
		} else {
			break
		}
	}
	return left +1, right -1
}

func main() {
	fmt.Println(longestPalindrome("ac"))
}
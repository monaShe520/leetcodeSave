package main

import "fmt"

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 示例 1：
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49 
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

func maxArea(height []int) int {
	maxLeft := 0
	maxRight := len(height)-1
	max := 0
	for {
		if maxLeft == maxRight {
			break
		}
		h := height[maxLeft]
		if height[maxLeft] > height[maxRight] {
			h = height[maxRight]
		}
		if (maxRight-maxLeft) * h > max {
			max = (maxRight-maxLeft) * h
		}
		if height[maxLeft] > height[maxRight] {
			maxRight  = maxRight - 1
		} else {
			maxLeft = maxLeft + 1
		}

	}
	
	return max
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
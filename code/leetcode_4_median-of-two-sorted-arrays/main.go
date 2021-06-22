package main

import "fmt"

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 示例 3：
// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000
// 示例 4：
// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000
// 示例 5：
// 输入：nums1 = [2], nums2 = []
// 输出：2.00000
// 提示：
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	index1 := 0
	index2 := 0
	sum := []int{}
	for {
		if index1 >= len(nums1) {
			sum = append(sum, nums2[index2:]...)
			break
		} else if index2 >= len(nums2) {
			sum = append(sum, nums1[index1:]...)
			break
		}
		if nums1[index1] <= nums2[index2] {
			sum = append(sum, nums1[index1])
			index1 = index1 + 1
		} else {
			sum = append(sum, nums2[index2])
			index2 = index2 + 1
		}
	}
	if len(sum) % 2 == 1 {
		return float64(sum[len(sum)/2])
	} else {
		return (float64(sum[len(sum)/2-1])+ float64(sum[len(sum)/2]))/2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
}
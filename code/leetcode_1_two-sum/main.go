package main

import "fmt"

import "sort"


func twoSum(nums []int, target int) []int {

	indexMap := make(map[int][]int, len(nums)) // 通过值反差下标
	for k, v := range nums {
		if indexMap[v] == nil {
			indexMap[v] = []int{}
		}
		indexMap[v] = append(indexMap[v], k)
	}
	sort.Ints(nums)
	for k := 0 ; k < len(nums) - 1; k++ {
		if 2 * nums[k] == target && len(indexMap[nums[k]]) == 2 {
			return indexMap[nums[k]]
		}
		if res, status := indexMap[target-nums[k]]; status {
			return []int{res[0], indexMap[nums[k]][0]}
		}
	}
	return []int{}
}

func main(){
	fmt.Println(twoSum([]int{3,3}, 6))
}
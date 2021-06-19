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
	vaildNums := []int{}
	for k := len(nums) - 1 ; k > 0; k-- {
		if nums[k] + nums[0] < target {
			vaildNums = nums[0:k+1]
			break
		} else if nums[k] + nums[0] == target {
			t1 := indexMap[nums[k]][len(indexMap[nums[k]])-1]
			if len(indexMap[nums[k]]) == 1 {
				indexMap[nums[k]] = []int{}
			} else {
				indexMap[nums[k]] = indexMap[nums[k]][0:len(indexMap[nums[k]])-1]
			}
			t2 := indexMap[nums[0]][0]
			return []int{t1, t2}	
		}
	}
	for k := 0 ; k < len(vaildNums) - 1; k++ {
		if 2 * vaildNums[k] == target && len(indexMap[vaildNums[k]]) == 2 {
			return indexMap[vaildNums[k]]
		}
		if res, status := indexMap[target-vaildNums[k]]; status {
			return []int{res[0], indexMap[vaildNums[k]][0]}
		}
	}
	return []int{}
}

func main(){
	fmt.Println(twoSum([]int{3,3}, 6))
}
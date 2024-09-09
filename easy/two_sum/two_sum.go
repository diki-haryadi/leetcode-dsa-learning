package main

func twoSum(nums []int, target int) []int {
	var result = []int{0, 0}
	var numToIndex = map[int]int{}

	for i, num := range nums {
		if index, exist := numToIndex[target-num]; exist {
			return []int{index, i}
		} else {
			numToIndex[num] = i
		}
	}
	return result
}

func twoSumVariant1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

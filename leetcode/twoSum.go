package main

import "fmt"

func twoSum(nums []int, target int) []int {

	//fin := []int{}

	for i := 0; i < len(nums); i++ {

		x := nums[i]

		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			if x+y == target {
				// fin = append(fin, i)
				// fin = append(fin, j)
				// break
				return []int{i, j}
			}

		}
	}
	return []int{}
}

func main() {

	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}

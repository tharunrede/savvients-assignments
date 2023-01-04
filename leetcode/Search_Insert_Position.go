package main

import (
	"fmt"
)

func main() {

	initial_arr := []int{1, 3, 5, 6}
	var target int
	fmt.Println("Array is:", initial_arr)

	fmt.Println("Enter any number to insert/ to see if it is in the array: ")
	fmt.Scanln(&target)
	fmt.Println(searchInsert(initial_arr, target))

}

func searchInsert(nums []int, target int) int {

	//flag := false
	position_inserted := -1

	for pos, val := range nums {
		//fmt.Println("pos, val, target", pos, val, target)

		if target == val {
			//fmt.Println("in tar== val")
			//flag = true
			position_inserted = pos
			break
		} else if target < val {
			//fmt.Println("in tar> val")
			nums = append(nums[:pos+1], nums[pos:]...)
			nums[pos] = target
			position_inserted = pos
			break

		} else if target > nums[len(nums)-1] {
			nums = append(nums, target)
			position_inserted = len(nums) - 1
			break
		}
	}

	return position_inserted
}

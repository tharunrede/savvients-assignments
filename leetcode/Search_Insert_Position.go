package main

import (
	"fmt"
)

func main() {

	initial_arr := []int{1, 2, 3, 4, 7, 8, 9}
	var target int
	fmt.Println("Array is:", initial_arr)

	fmt.Println("Enter any number to insert/ to see if it is in the array: ")
	fmt.Scanln(&target)

	flag := false
	position_inserted := -1

	for pos, val := range initial_arr {
		//fmt.Println("pos, val, target", pos, val, target)

		if target == val {
			//fmt.Println("in tar== val")
			flag = true
			position_inserted = pos
			break
		} else if target < val {
			//fmt.Println("in tar> val")
			initial_arr = append(initial_arr[:pos+1], initial_arr[pos:]...) // index < len(a)
			initial_arr[pos] = target
			position_inserted = pos
			break

		} else if target > initial_arr[len(initial_arr)-1] {
			initial_arr = append(initial_arr, target)
			position_inserted = len(initial_arr) - 1
			break
		}
	}

	if flag == false {
		fmt.Println("Element not found but inserted at ", position_inserted, "index")
	} else {
		fmt.Println("Element found at", position_inserted, "index")
	}

	fmt.Println("Final initial_arr", initial_arr)
}

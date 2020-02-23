package main

import "fmt"

func sum(nums []int, index int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[index] + sum(nums[(index+1):], 0)
}

func main() {

	var nums = []int{10, 20, 30, 20, 50, 100, 200}

	res := sum(nums, 0)

	fmt.Println(res)

}

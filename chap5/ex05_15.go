package main

import "fmt"

func minMax(nums ...int) (int, int, error) {
	if len(nums) == 0 { return -1, -1, fmt.Errorf("no number is provided") }
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num < min { min = num }
		if num > max { max = num }
	}

	return min, max, nil
}

func minMax2(num int, nums ...int) (min int, max int) {
	min = num
	max = num
	for _, item := range nums {
		if item < min { min = item }
		if item > max { max = item }
	}
	return
}

func main() {
	fmt.Println(minMax(1,23,4,5))
	fmt.Println(minMax2(1,23,4,5))
	fmt.Println(minMax())
}

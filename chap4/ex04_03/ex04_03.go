package main

import "fmt"

const ArrSize = 10

func reverse(ptr *[ArrSize]int) {
	for i, j := 0, ArrSize - 1; i < j;  i, j = i + 1, j - 1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	nums := [ArrSize]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)
	reverse(&nums)
	fmt.Println(nums)
}

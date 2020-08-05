package main

import "fmt"

func rotateRight(s []int, rotateBy int) {
	n := len(s)
	var period int
	for i:= 0; period == 0 || i < period; i++ {
		idx := i
		temp := s[i]
		count := 0
		for {
			count += 1
			nextIdx := (idx + n - rotateBy) % n
			if nextIdx == i { break }
			s[idx] = s[nextIdx]
			idx = nextIdx
		}
		s[idx] = temp
		if period == 0 { period = n / count }
	}
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl)
	rotateRight(sl, 5)
	fmt.Println(sl)
}

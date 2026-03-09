package main

import "fmt"

func main() {

	nums := make([]int, 0, 50)
	prevCap := cap(nums)

	for i := 0; i < 50; i++ {
		nums = append(nums, i)

		if cap(nums) != prevCap {
			fmt.Printf("len=%d cap=%d\n", len(nums), cap(nums))
			prevCap = cap(nums)
		}
	}
}

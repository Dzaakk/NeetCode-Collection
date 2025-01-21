package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mapNums := make(map[int]bool)
	for _, n := range nums {
		if mapNums[n] {
			return false
		}
		mapNums[n] = true
	}
	return true
}
func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 3, 4, 1}
	fmt.Println(containsDuplicate(nums1))
	fmt.Println(containsDuplicate(nums2))
}

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}

	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 3, 3, 3}
	k1 := 2
	fmt.Println(topKFrequent(nums1, k1))

	nums2 := []int{7, 7}
	k2 := 1
	fmt.Println(topKFrequent(nums2, k2))
}

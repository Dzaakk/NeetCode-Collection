package main

// O(n2)
// func twoSum(nums []int, target int) [2]int {
// 	idx := [2]int{}
// 	for i := 0; i < len(nums); i++ {
// 		a := nums[i]
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			b := nums[j]

// 			if a+b == target {
// 				idx[0] = i
// 				idx[1] = j

// 				return idx
// 			}
// 		}

// 	}
// 	return idx
// }

// better approach O(n)
func twoSum(nums []int, target int) [2]int {
	idxMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := idxMap[complement]; found {
			return [2]int{j, i}
		}
		idxMap[num] = i
	}

	return [2]int{}
}

// func main() {
// 	nums1 := []int{3, 4, 5, 6}
// 	target1 := 7
// 	nums2 := []int{4, 5, 6}
// 	target2 := 10

// 	fmt.Println(twoSum(nums1, target1))
// 	fmt.Println(twoSum(nums2, target2))
// }

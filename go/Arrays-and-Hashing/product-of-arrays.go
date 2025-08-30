package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}
	// [1,2,3,4]
	// fmt.Printf("NUMS = %d\n\n", nums)
	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		// fmt.Printf("prefix = %d\n", prefix)
		// fmt.Printf("prefix = prefix * nums[%d]\nprefix = %d * %d\n", i, prefix, nums[i])
		prefix *= nums[i]
		// fmt.Printf("RES = %d\nnew prefix = %d\n\n", res, prefix)
	}
	// fmt.Println("-------------------------------")
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		// fmt.Printf("postfix = %d\n", postfix)
		// fmt.Printf("postfix = postfix * nums[%d]\npostfix = %d * %d\n", i, postfix, nums[i])
		postfix *= nums[i]
		// fmt.Printf("RES = %d\nnew postfix= %d\n\n", res, postfix)
	}

	return res
}

// func main() {
// 	nums := []int{1, 2, 3, 4}

// 	fmt.Println(productExceptSelf(nums))
// }

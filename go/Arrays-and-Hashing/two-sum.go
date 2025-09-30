package main

//11.40
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		dif := target - v

		if j, ok := m[dif]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	return []int{0, 0}
}

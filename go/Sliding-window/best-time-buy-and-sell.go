package main

// import "fmt"
// 17
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] > prices[r] {
			l++
			continue
		} else {
			maxP = max(maxP, prices[r]-prices[l])
		}
		r++
	}

	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	prices1 := []int{10, 1, 5, 6, 7, 1} // 6
// 	prices2 := []int{10, 8, 7, 5, 2}    // 0 no profit

// 	fmt.Println(maxProfit(prices1))
// 	fmt.Println(maxProfit(prices2))
// }

package main

// import "fmt"

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxP {
				maxP = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxP
}

// func main() {
// 	prices1 := []int{10, 1, 5, 6, 7, 1} // 6
// 	prices2 := []int{10, 8, 7, 5, 2}    // 0 no profit

// 	fmt.Println(maxProfit(prices1))
// 	fmt.Println(maxProfit(prices2))
// }

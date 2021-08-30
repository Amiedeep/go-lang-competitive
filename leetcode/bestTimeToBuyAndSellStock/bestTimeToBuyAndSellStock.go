package main

import "fmt"

func main() {

	prices := []int{7,6,4,3,1}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
    profit := 0

	min := prices[0]

	for  i:=1; i< len(prices); i++ {
		if (prices[i] < min) {
			min = prices[i]
			continue
		}
		if (prices[i] - min > profit) {
			profit = prices[i] - min
		}
	}

	return profit
}
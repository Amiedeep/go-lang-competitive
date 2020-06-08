package main

import "fmt"

var maxPrice int

func findMaxPrice(length int, prices []int, currentPrice int) {
	if length < 0 {
		return
	}
	if currentPrice > maxPrice {
		maxPrice = currentPrice
	}

	for i := 0; i < len(prices); i++ {
		findMaxPrice(length-(i+1), prices, currentPrice+prices[i])
	}
}

func main() {

	length := 8
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20}
	// prices := []int{3, 5, 8, 9, 10, 17, 17, 20}

	findMaxPrice(length, prices, 0)

	fmt.Println(maxPrice)
}

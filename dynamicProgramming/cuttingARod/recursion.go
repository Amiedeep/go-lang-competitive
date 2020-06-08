package main

import "fmt"

var maxPrice int

func findMaxPrice(length int, prices []int, index int, currentPrice int) {
	if length < 0 {
		return
	}
	if currentPrice > maxPrice {
		maxPrice = currentPrice
	}

	for i := index; i < len(prices); i++ {
		findMaxPrice(length-(i+1), prices, i+1, currentPrice+prices[i])
		findMaxPrice(length-(i+1), prices, i, currentPrice+prices[i])
		findMaxPrice(length, prices, i+1, currentPrice)
	}
}

func main() {

	length := 8
	// prices := []int{1, 5, 8, 9, 10, 17, 17, 20}
	prices := []int{3, 5, 8, 9, 10, 17, 17, 20}

	findMaxPrice(length, prices, 0, 0)

	fmt.Println(maxPrice)
}

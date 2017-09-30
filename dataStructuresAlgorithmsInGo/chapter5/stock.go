package main

import "fmt"

// “iven a list, in which nth element is the price of the stock on nth day. You are asked to buy once and sell once, on what date you will be buying and at what date you will be selling to get maximum profit.
// Or
// Given a list of numbers, you need to maximize the difference between two numbers, such that you can subtract the number, which appear before form the number that appear after it.
//
//
// First approach:  Brute force, for each element in list find if there is some other element whose difference is maximum. This is done using two for loop, first loop to select, buy date index and the second loop to find its selling date entry.
//
// The Time Complexity is O(n2) and Space Complexity is O(1)”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// “Second approach: Another clever solution is to keep track of the smallest value seen so far from the start. At each point, we can find the difference and keep track of the maximum profit. This is a linear solution.
// The Time Complexity of the algorithm is O(n) time. Space Complexity for creating count list is also O(1)”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func maxProfit(stocks []int) {
	size := len(stocks)
	var buy, sell, currMin, currProfit, maxProfit int

	for i := 0; i < size; i++ {
		if stocks[i] < stocks[currMin] {
			currMin = i
		}
		currProfit = stocks[i] - stocks[currMin]
		if currProfit > maxProfit {
			buy = currMin
			sell = i
			maxProfit = currProfit
		}
	}
	fmt.Println("Purchase day is-", buy, "at price", stocks[buy])
	fmt.Println("Sell day is-", sell, "at price", stocks[sell])
	fmt.Println("Max Profit ::", maxProfit)
}

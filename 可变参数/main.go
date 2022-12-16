package main

import "fmt"

func main() {
	getSum(1, 2, 9, 9)
}

func getSum(start int, nums ...int) {
	sums := start
	for i := 0; i < len(nums); i++ {
		sums += nums[i]
	}
	fmt.Println(sums)
}

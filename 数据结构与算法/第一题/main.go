package main

import "fmt"

func main() {
	target := 6
	arr := []int{3, 2, 4, 9, 8, 7, 44, 99, 789}
	v1 := twoSum(arr, target)
	fmt.Println(v1)

}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		//if p, ok := hashTable[target-x]; ok {
		//	return []int{p, i}
		//}
		hashTable[x] = i
		fmt.Println(hashTable)
	}
	return nil
}

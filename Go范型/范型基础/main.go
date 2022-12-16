package main

import "fmt"

func main() {
	str := []string{"aaaa", "cccc"}
	myint := []int{3, 2, 1}
	PrintList(str)
	PrintList(myint)
}

func PrintList[T comparable](arr []T) {
	for _, i := range arr {
		fmt.Println(i)
	}
}

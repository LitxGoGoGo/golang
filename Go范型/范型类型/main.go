package main

import "fmt"

func main() {
	type Slice[T int | float32 | float64] []T
	var a Slice[int] = []int{4, 5, 6}
	for b := range a {
		fmt.Println(b)
	}

	type MyMap[KEY int | string, VALUE float64 | float32] map[KEY]VALUE

	var m1 MyMap[string, float32] = map[string]float32{
		"GO":   9.9,
		"JAVA": 10,
	}
	fmt.Println(m1)
}

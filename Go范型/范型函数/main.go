package main

import "fmt"

type MySlice[T int | float64] []T

func main() {
	var test MySlice[int] = []int{1, 2, 3}
	fmt.Println(test.Sum())
	var f MySlice[float64] = []float64{1.0, 2.2, 3.5}
	fmt.Println(f.Sum())
	fmt.Println(Add("asd", "1231"))
}

func (s MySlice[T]) Sum() T {
	var sum T
	for _, i := range s {
		sum += i
	}
	return sum
}

// Add 范型函数
func Add[T string | int | float64](a T, b T) T {
	return a + b
}

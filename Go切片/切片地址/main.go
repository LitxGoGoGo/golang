package main

import "fmt"

func appendSliceInPlace(slice []int, data ...int) {
	// 这里相当于slice的二级指针，因为slice本身就是一个指针，
	// 所以这里我们让可以让slice指向一个新的数组[1 2 3 4 5 6]
	// 注意：append 其实会返回一个新的数组
	slice = append(slice, data...)
	fmt.Println("1", slice)
}

func main() {
	slice := []int{1, 2, 3}
	appendSliceInPlace(slice, 4, 5, 6)
	fmt.Println("2", slice)
}

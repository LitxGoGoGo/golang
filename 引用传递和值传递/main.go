package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	update(arr)
	fmt.Println(arr)
	fmt.Printf("地址1是%p", &arr)

}

func update(arr2 []int) {
	fmt.Println("前", arr2)
	arr2[0] = 99
	fmt.Println("后", arr2)
	fmt.Printf("地址是%p", &arr2)
}

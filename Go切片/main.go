package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	var newslice []int
	v1 := append(newslice, slice...)
	fmt.Printf("%p \n", newslice)

	fmt.Printf("%p \n", v1)
	fmt.Println(v1)
	v2 := append(newslice, slice...)
	fmt.Printf("%p \n", newslice)

	fmt.Printf("%p \n", v2)
	fmt.Println(v2)

	months := []string{1: "January"}
	fmt.Println("xxxxxxx")
	fmt.Printf("%p,\n", &months)
	fmt.Printf("%p,\n", months)
	fmt.Printf("%p,\n", &months[0])

	fmt.Printf("%v \n", &months)
	fmt.Printf("%v \n", months)
	fmt.Printf("%v \n", &months[0])
	//newSlice := slice[0:3]
	//fmt.Printf("%T\n", slice)
	//fmt.Println("before modifying underlying array:")
	//fmt.Println("slice: ", slice)
	//fmt.Println("newSlice: ", newSlice)
	//fmt.Println()
	//
	//newSlice[0] = 6
	//fmt.Println("after modifying underlying array:")
	//fmt.Println("slice: ", slice)
	//fmt.Println("newSlice: ", newSlice)
}

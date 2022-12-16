package main

import "fmt"

func main() {
	num1 := 0
	v1 := &num1
	v2 := &num1
	var v3 *int
	v3 = v2
	*v3 = 100
	fmt.Println(num1, &v1, &v2, &v3)
	fmt.Println(v1 == v2)
	num1 = 10
	fmt.Println(*v1, *v2)
}

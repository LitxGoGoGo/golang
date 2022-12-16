package main

import "fmt"

type int8AAAA int8

type MyInt interface {
	int | ~int8 | int16
}

func GetMaX[T MyInt](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var a int8AAAA = 10
	var b int8AAAA = 20
	fmt.Println(GetMaX(a, b))
}

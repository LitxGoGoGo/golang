package main

import "fmt"

func FuncSlice(s []int, t int) {
	fmt.Println(s)
	fmt.Printf("1%p \n", s)
	s[0]++
	fmt.Println(s)
	fmt.Printf("2%p \n", s)
	s = append(s, t)
	fmt.Println(s)
	fmt.Printf("3%p \n", s)
	s[0]++
	fmt.Println(s)
	fmt.Printf("4%p \n", s)
}
func main() {
	a := []int{0, 1, 2, 3}
	fmt.Printf("5%p \n", a)
	fmt.Printf("55%p", &a)
	FuncSlice(a, 4)
	fmt.Printf("6%p \n", a)
	fmt.Println(a)
}

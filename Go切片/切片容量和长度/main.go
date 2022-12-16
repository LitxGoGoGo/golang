package main

import "fmt"

func main() {
	//s := []int{2, 3, 5, 7, 11, 13}
	//bianli(s)
	//printSlice(s)
	//
	//// 截取切片使其长度为 0
	//s = s[:0]
	//bianli(s)
	//printSlice(s)
	//
	//// 拓展其长度
	//s = s[:4]
	//bianli(s)
	//printSlice(s)
	//
	//// 舍弃前两个值
	//s = s[2:]
	//bianli(s)
	//printSlice(s)
	//
	//// 再次取回之前的值
	//s = s[:4]
	//bianli(s)
	//printSlice(s)
	a := []int{1, 2, 3}
	printSlice(a)
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	printSlice(a)
	a = a[:1]
	printSlice(a)

	//a = append(a[:0], a[N:]...) // 删除开头N个元素
	//printSlice(a)

	d := []byte{'r', 'o', 'a', 'd'}
	d = d[2:]
	// e == []byte{'a', 'd'}
	d[1] = 'm'

	for _, i := range d {
		fmt.Printf("%T \n", i)
		fmt.Println(string(i))
	}
	fmt.Printf("%v \n", d)
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v 类型为=%T\n", len(s), cap(s), s, s)
}

func bianli(s []int) {
	for _, i := range s {
		fmt.Println(i)
	}
}

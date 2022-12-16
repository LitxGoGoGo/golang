package main

import "fmt"

type Student struct {
	name string
}

type Person interface {
	sayHello(name string) string
}

func main() {
	//var s Person = &Student{name: "3215161"}
	s := &Student{name: "4894651"}
	//v, ok := s.(Person)
	fmt.Println(s.sayHello("litianxiang"))
	//if !ok {
	//	fmt.Printf("%v\n", v)
	//} else {
	//	fmt.Println("hhhh")
	//}
}

func (s Student) sayHello(name string) string {
	return fmt.Sprintf("%v: Hello %v, nice to meet you.\n", s.name, name)
}

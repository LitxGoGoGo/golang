package main

import "fmt"

type USB interface {
	write()
	read()
}


type Phone struct {
	name string
}

type Computer struct {
	Phone
}



func main() {
	p := Phone{name: func(s string) string {
		return "手机" + s
	}("哈哈哈")}
	p.read()

	c := Computer{
		Phone{name: "woshi电脑"},
	}
	c.read()
}

func (p *Phone) write() {
	fmt.Println(p.name, "这是手机的write方法")
}

func (p *Phone) read() {
	fmt.Println(p.name, "这是手机的read方法")
}




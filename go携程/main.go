package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}
	time.Sleep(3 * time.Second)
}

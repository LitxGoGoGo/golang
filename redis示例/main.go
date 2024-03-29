package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.100.1:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("连接成功。", pong)

	key := "key1"
	//err = client.Set(key, "987987dasd", 0).Err()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("写入成功。")

	value, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("读取成功：" + value)
}

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

var T struct {
	id int
	a  int
	b  int
}

func main() {
	//配置MySQL连接参数
	username := "root"       //账号
	password := "123456"     //密码
	host := "121.37.248.123" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "test"         //数据库名
	timeout := "10s"         //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		panic("连接数据库失败")
	}

	db.AutoMigrate(&Product{})
	//var product Product
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//db.First(&product, 1) // 找到id为1的产品
	//fmt.Println(&product)

	// 增
	//db.Create(&Product{Code: "L1212", Price: 7000})
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {

		go func(a int) {
			defer wg.Done()
			db.Create(&Product{Code: "L1212", Price: 1000 + a})
		}(i)
	}
	wg.Wait()
}

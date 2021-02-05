package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float32
}

func GetProduct(id int) *Product {
	r := rand.Intn(10)
	if r < 6 {
		time.Sleep(5 * time.Second)
	}
	return &Product{
		ID:    id,
		Name:  "彩色珊瑚",
		Price: 9.01,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		// hystrix 设置command配置文件
		configA := hystrix.CommandConfig{
			Timeout: 4000,
			//MaxConcurrentRequests:  0,
			//RequestVolumeThreshold: 0,
			//SleepWindow:            0,
			//ErrorPercentThreshold:  0,
		}
		// hystrix 设置Configure文件
		hystrix.ConfigureCommand("get_product", configA)

		err := hystrix.Do("get_product", func() error {
			product := GetProduct(1)
			fmt.Println(product)
			return nil
		}, nil)

		if err != nil {
			fmt.Println("发生错误了")
			fmt.Println(err)
		}
	}

}

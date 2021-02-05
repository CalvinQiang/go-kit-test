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

func RecProduct() (Product, error) {
	return Product{
		ID:    999,
		Name:  "推荐商品",
		Price: 120,
	}, nil
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
		getProductChan := make(chan Product, 1)
		errs := hystrix.Go("get_product", func() error {
			product := GetProduct(1)
			getProductChan <- *product
			return nil
		}, func(e error) error {
			product, err := RecProduct()
			getProductChan <- product
			return err
		})

		select {
		case getProduct := <-getProductChan:

			fmt.Println(getProduct)
		case errors := <-errs:
			fmt.Println("发生错误了")
			fmt.Println(errors)

		}
	}

}

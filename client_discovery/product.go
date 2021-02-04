package main

import (
	endpointClient "com.calvin.client_discovery/endpoint"
	"com.calvin.client_discovery/transport"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httpTransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"net/url"
	"os"
	"time"
)

func main() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.0.106:8500" // 注册中心地址
	apiClient, _ := consulapi.NewClient(config)
	client := consul.NewClient(apiClient)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	tags := []string{"primary"}
	// 实时查询服务实例的状态信息
	instancer := consul.NewInstancer(client, logger, "userservice", tags, true)

	{
		f := func(serviceUrl string) (endpoint.Endpoint, io.Closer, error) {
			tart, _ := url.Parse("http://" + serviceUrl)
			return httpTransport.NewClient("GET", tart, transport.GetUserInfoRequest, transport.GetUserInfoResponse).Endpoint(), nil, nil
		}
		endpointer := sd.NewEndpointer(instancer, f, logger)
		// 轮询获取服务
		robin := lb.NewRandom(endpointer, int64(time.Now().Nanosecond()))
		for {
			time.Sleep(500 * time.Millisecond)
			getUserInfo, _ := robin.Endpoint()
			ctx := context.Background()
			res, err := getUserInfo(ctx, endpointClient.UserRequest{Uid: 101})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			userInfo := res.(endpointClient.UserResponse)
			fmt.Println(userInfo)
		}
	}
}

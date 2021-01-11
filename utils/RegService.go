package utils

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var client *consulapi.Client

// 服务器节点地址
const nodeAddress = "192.168.0.110"

// 注册中心地址
const rcAddress = "192.168.0.110"

func init() {
	config := consulapi.DefaultConfig()
	config.Address = rcAddress + ":8500"
	// 新创建一个客户端请求
	newClient, err := consulapi.NewClient(config)
	if err != nil {
		// 直接抛出异常
		log.Fatal(err)
	}
	client = newClient
}
func RegService() {
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "userservice"
	reg.Name = "userservice"
	reg.Address = nodeAddress
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://" + nodeAddress + ":8080/health"

	reg.Check = &check

	err := client.Agent().ServiceRegister(&reg)
	if err != nil {
		// 直接抛出异常
		log.Fatal(err)
	}
}

func UnRegService() {
	_ = client.Agent().ServiceDeregister("userservice")
}

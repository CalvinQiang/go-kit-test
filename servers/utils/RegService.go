package utils

import (
	consulapi "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"log"
	"strconv"
)

var client *consulapi.Client

// 服务器节点地址
const nodeAddress = "192.168.0.102"

// 注册中心地址
const rcAddress = "192.168.0.102"

// 服务名称
var ServiceName string

// 服务端口号
var ServicePort int

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
	if ServiceName == "" {
		ServiceName = "userservice"
	}
}

func SetServiceNameAndPort(serviceName string, servicePort int) {
	ServiceName = serviceName
	ServicePort = servicePort
}

func RegService() {
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = uuid.NewV1().String()
	reg.Name = ServiceName
	reg.Address = nodeAddress
	reg.Port = ServicePort
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://" + nodeAddress + ":" + strconv.Itoa(ServicePort) + "/health"

	reg.Check = &check

	err := client.Agent().ServiceRegister(&reg)
	if err != nil {
		// 直接抛出异常
		log.Fatal(err)
	}
}

func UnRegService() {
	_ = client.Agent().ServiceDeregister(ServiceName)
}

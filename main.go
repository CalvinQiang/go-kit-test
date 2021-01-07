package main

import (
	"com.calvin.service/endpoint"
	"com.calvin.service/service"
	"com.calvin.service/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main() {
	user := service.UserService{}
	endPoint := endpoint.GetUserEndPoint(user)

	// construct a new server， implements http.handler and wrap the endpoint
	serverHandler := httpTransport.NewServer(endPoint, transport.DecodeUserRequest, transport.EncodeUserResponse)
	// 监听端口，并且使用serverHandler处理随之而来的请求
	_ = http.ListenAndServe(":8080", serverHandler)
}

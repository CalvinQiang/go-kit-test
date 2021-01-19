package main

import (
	"com.calvin.service/endpoint"
	"com.calvin.service/service"
	"com.calvin.service/transport"
	"com.calvin.service/utils"
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	goMux "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	user := service.UserService{}
	endPoint := endpoint.GetUserEndPoint(user)

	// construct a new server， implements http.handler and wrap the endpoint
	serverHandler := httpTransport.NewServer(endPoint, transport.DecodeUserRequest, transport.EncodeUserResponse)
	r := goMux.NewRouter()                                                    // 这里我们引入了第三方路由
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler) // 设置path格式的路由
	r.Methods("GEt").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go (func() {
		// 注册服务
		utils.RegService()
		// 监听端口，并且使用serverHandler处理随之而来的请求
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Println("启动出错，退出")
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		singC := make(chan os.Signal)
		signal.Notify(singC, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-singC)
	})()

	// 等待信号量
	getErr := <-errChan
	utils.UnRegService()
	log.Println("接收到信号量，退出")
	log.Println(getErr)
}
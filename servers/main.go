package main

import (
	"com.calvin.service/endpoint"
	"com.calvin.service/service"
	"com.calvin.service/transport"
	"com.calvin.service/utils"
	"flag"
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	goMux "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	// 获取启动参数
	serviceName := flag.String("name", "", "服务的名字")
	servicePort := flag.Int("port", 8001, "服务的端口号")
	flag.Parse()
	fmt.Println(*serviceName)
	if *serviceName == "" {
		log.Fatal("未指定服务的名字")
	}
	user := service.UserService{}
	endPoint := endpoint.GetUserEndPoint(user)

	// construct a new server， implements http.handler and wrap the endpoint
	serverHandler := httpTransport.NewServer(endPoint, transport.DecodeUserRequest, transport.EncodeUserResponse)
	r := goMux.NewRouter()                                                    // 这里我们引入了第三方路由
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler) // 设置path格式的路由
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go (func() {
		// 设置服务名和端口号码
		utils.SetServiceNameAndPort(*serviceName, *servicePort)
		// 注册服务
		utils.RegService()
		// 监听端口，并且使用serverHandler处理随之而来的请求
		err := http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
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

package main

import (
	"com.calvin.service/endpoint"
	"com.calvin.service/service"
	"com.calvin.service/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	goMux "github.com/gorilla/mux"
	"net/http"
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
	// 监听端口，并且使用serverHandler处理随之而来的请求
	_ = http.ListenAndServe(":8080", r)
}

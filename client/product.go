package main

import (
	"com.calvin.client/endpoint"
	"com.calvin.client/transport"
	"context"
	"fmt"
	transportHttp "github.com/go-kit/kit/transport/http"
	"net/url"
	"os"
)

func main() {
	target, _ := url.Parse("http://localhost:8080")
	client := transportHttp.NewClient("GET", target, transport.GetUserInfoRequest, transport.GetUserInfoResponse)
	getUserInfo := client.Endpoint()

	ctx := context.Background()
	res, err := getUserInfo(ctx, endpoint.UserRequest{Uid: 101})
	if err != nil {
		os.Exit(1)
	}

	userInfo := res.(endpoint.UserResponse)
	fmt.Println(userInfo)

}

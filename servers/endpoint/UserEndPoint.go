package endpoint

import (
	"com.calvin.service/service"
	"com.calvin.service/utils"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"strconv"
)

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}

func GetUserEndPoint(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		if r.Method == "GET" {
			result := userService.GetName(r.Uid) + strconv.Itoa(utils.ServicePort)
			return UserResponse{Result: result}, nil
		} else if r.Method == "DELETE" {
			result := userService.DelUser(r.Uid)
			return UserResponse{Result: result}, nil
		} else {
			return nil, errors.New("未知方法，无法处理")
		}
	}
}

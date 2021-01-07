package endpoint

import (
	"com.calvin.service/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GetUserEndPoint(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := userService.GetName(r.Uid)
		return UserResponse{Result: result}, nil
	}
}

package transport

import (
	"com.calvin.service/endpoint"
	"context"
	"encoding/json"
	"errors"
	goMux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := goMux.Vars(r) // 使用vars方法，提取路由上的参数
	if result, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(result)
		return endpoint.UserRequest{Uid: uid}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

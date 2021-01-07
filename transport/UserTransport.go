package transport

import (
	"com.calvin.service/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// 限定参数来源
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return endpoint.UserRequest{Uid: uid}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

package transport

import (
	"com.calvin.service/endpoint"
	"com.calvin.service/utils"
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
		return endpoint.UserRequest{Uid: uid, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("content-type", contentType)
	if myError, ok := err.(*utils.MyError); ok {
		w.WriteHeader(myError.Code)
		_, _ = w.Write(body)
	} else {
		w.WriteHeader(404)
		_, _ = w.Write(body)
	}

}

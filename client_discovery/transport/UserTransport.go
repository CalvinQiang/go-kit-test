package transport

import (
	endpointClient "com.calvin.client_discovery/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetUserInfoRequest(c context.Context, req *http.Request, r interface{}) error {
	userRequest := r.(endpointClient.UserRequest)
	req.URL.Path += "/user/" + strconv.Itoa(userRequest.Uid)
	return nil
}

func GetUserInfoResponse(c context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode > 400 {
		return nil, errors.New("no data")
	}

	var userResponse endpointClient.UserResponse
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, err
}

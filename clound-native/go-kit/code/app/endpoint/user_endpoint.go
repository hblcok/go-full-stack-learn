package endpoint

import (
	"context"

	"hb.study/clound-native/go-kit/code/app/service"

	"github.com/go-kit/kit/endpoint"
)

// 请求
type UserRequest struct {
	UserID int `json:"id"`
}

// 响应
type UserResponse struct {
	UserName string `json:"userName"`
}

// 生成endpoint
func GenUserEndpoint(service service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		name := service.GetUserName(r.UserID)
		return UserResponse{
			UserName: name,
		}, nil
	}
}

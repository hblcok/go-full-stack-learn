package main

import (
	"context"
	"fmt"
	"go-full-stack-learn/clound-native/grpc/code/demo_product/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 1.创建连接套接字，端口和服务端一致，并且先不使用HTTPS(即不安全调用)
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("客户端连接服务端异常，err:", err)
	}
	defer conn.Close()

	// 2.创建grpc客户端服务，调用 proto 生成的代码中的方法.
	prodClient := service.NewProdServiceClient(conn)

	// 3.调用远程服务，并处理响应
	prodResp, err := prodClient.GetProdName(context.Background(), &service.ProdRequest{ProdID: 10})
	if err != nil {
		log.Fatal("调用商品服务异常，err:", err)
	}

	fmt.Println("商品名称为：", prodResp.ProdName)
}

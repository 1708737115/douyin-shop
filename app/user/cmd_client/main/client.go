package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user/userservice"
)

func main() {
    // 创建客户端，指定服务地址
    cli, err := userservice.NewClient("user_service", client.WithHostPorts("localhost:8007"))
    if err != nil {
        fmt.Println("Error creating client:", err)
        return
    }

    // 调用服务的方法
    req := &user.LoginReq{
        // 根据实际的请求结构体填写
		Email: string("test@example.com"),
		Password: string("password"),
    }
    res, err := cli.Login(context.Background(), req)
    if err != nil {
        fmt.Println("Error calling method:", err)
        return
    }
    fmt.Printf("Response: %v\n", res)
}
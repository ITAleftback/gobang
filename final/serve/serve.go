package main

import (

	"final/proto"
	"final/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)


func main() {


	////========================================================================================


	////监听所有网卡8081端口的TCP连接
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	fmt.Println("已开始监听")
	s := grpc.NewServer() //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	proto.RegisterRegisterServer(s,&service.Server{})
	proto.RegisterLoginServer(s,&service.Server{})
	proto.RegisterReadyServer(s,&service.Server{})
	proto.RegisterPlaychessServer(s,&service.Server{})
	proto.RegisterSurrenderServer(s,&service.Server{})

	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)

	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
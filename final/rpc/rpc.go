package rpc

import (
	"context"
	"final/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

// 建立RPC 注册服务
func RPCRegister(username string,password string){
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect: %v", err)
		return
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Register服务的客户端
	c :=proto.NewRegisterClient(conn)
	//注册
	if len(os.Args)>1{
		//经过学长提醒，我仔细去网上搜索了os.Args
		//传入的参数0时，返回的是文件路径
		//1和2都是输入的参数
		username = os.Args[1]
		password = os.Args[2]
	}
	//调用服务
	r,err:=c.Register(context.Background(),&proto.Register{Username:username,Password:password})
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(r.GetMessage())
}

// 建立RPC 登录服务
func RPCLogin(username string,password string){
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Login服务的客户端
	c :=proto.NewLoginClient(conn)
	//注册
	//同上面注释
	if len(os.Args)>1{
		username = os.Args[1]
		password = os.Args[1]
	}
	//调用服务
	r,err:=c.Login(context.Background(),&proto.Login{Username:username,Password:password})
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(r.GetMessage())
}
//建立RPC 准备服务
func RPCReady(){
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Ready服务的客户端
	c :=proto.NewReadyClient(conn)

	//调用服务
	r,err:=c.Ready(context.Background(),&proto.Ready{})
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(r.GetMessage())
}
//RPC的下棋服务
func RPCplaychess(i int32,j int32){
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Login服务的客户端
	c :=proto.NewPlaychessClient(conn)
	//调用服务
	r,err:=c.Playchess(context.Background(),&proto.Playchess{I:i,J:j})
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(r.GetMessage())
}
//认输
func RPCSurrender(){
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Peace服务的客户端
	c :=proto.NewSurrenderClient(conn)
	//调用服务
	r,err:=c.Surrender(context.Background(),&proto.Surrender{})
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(r.GetMessage())
}

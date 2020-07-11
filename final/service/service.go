package service

import (
	"context"
	"final/db"
	. "final/game/judge"
	. "final/game/play"
	"log"

	"final/proto"
)

type Server struct {}
//服务器端，注册的操作，玩家注册时，服务端将玩家注册的账号密码添加在数据库，sql语句在另一个层面
func (s *Server)Register(ctx context.Context, in *proto.Register) (*proto.Reply, error)  {
	log.Printf("Received:%v\n",in.GetUsername(),in.GetPassword())
	username:=in.GetUsername()
	password:=in.GetPassword()
	message:=db.Reg(username,password)

	return &proto.Reply{Message: message},nil
}
//服务器端登录的操作，玩家尝试登录时，服务端去匹配账号密码，并将成功与否的结果返回给玩家
func (s *Server) Login(ctx context.Context, in *proto.Login) (*proto.Reply, error){
	username:=in.GetUsername()
	password:=in.GetPassword()
	message:=db.Login(username,password)
	return &proto.Reply{Message: message},nil
}
//服务器端，玩家点击准备后的操作需要返回已准备的消息
func (s *Server)Ready(ctx context.Context, in *proto.Ready) (*proto.Reply, error)  {

	message:=db.Rea()
	return &proto.Reply{Message: message},nil
}
//获得玩家下棋的位置，并将之在服务端的棋盘下棋
func (s *Server)Playchess(ctx context.Context, in *proto.Playchess) (*proto.Reply, error)  {
	i:=in.GetI()
	j:=in.GetJ()
	message:=Play(i,j)
	return &proto.Reply{Message: message},nil
}
//给玩家的一个认输选择
func (s *Server)Surrender(ctx context.Context, in *proto.Surrender) (*proto.Reply, error)  {

	message:=Sua()
	return &proto.Reply{Message: message},nil
}


////////////////////////////=============================





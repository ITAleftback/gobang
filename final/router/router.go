package router

import (
	"final/rpc"
	"fmt"
	"net/http"
)

//注册的路由
//注册
func Reg(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	rpc.RPCRegister(username,password)
}
//登录
func Login(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	rpc.RPCLogin(username,password)
}
//准备
func Ready(w http.ResponseWriter,r *http.Request){
	rpc.RPCReady()
}
//下棋
func Playchess(w http.ResponseWriter,r *http.Request)  {
	var i,j int32
	//从键盘输入值,  i代表行  j代表列
	fmt.Println("请输入想下的位置")
	_, _ = fmt.Scanln(&i, &j)
	rpc.RPCplaychess(i,j)
}
//认输
func Surrender(w http.ResponseWriter,r *http.Request)  {
	rpc.RPCSurrender()
}

package main

import (
	"final/router"
	"log"
	"net/http"
)

func main()  {

	//注册路由
	http.HandleFunc("/Register", router.Reg)//注册
	http.HandleFunc("/Login", router.Login)//登录
	http.HandleFunc("/Ready",router.Ready)//准备  一开始默认为为准备状态
	http.HandleFunc("/Playchess",router.Playchess)//下棋
	http.HandleFunc("/Surrender",router.Surrender)//认输


	//http.HandleFunc("/Surrender",Surrenderp)
	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


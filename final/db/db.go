package db

import (
	. "final/db/init"
	."final/game/chessboard"
	"fmt"
	"log"
)


//有关准备的操作
func Rea()(message string){
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			fmt.Printf("%d\t", Board[i][j])

		}
		fmt.Println("\n")
	}
	message=`已准备`
	return
}
//这是 有关注册的数据库操作
func Reg(username string,password string)(message string){
	//注册sql语句
	stmt,err:= Db.Prepare(
		"insert into user(username,password)values(?,?)")
	if err!=nil{
		log.Fatal(err)
	}
	stmt.Exec(username,password)
	message=`注册成功`+username
	return
}
//有关登录的数据库操作
func Login(username string,password string)(message string){
	rows, err :=Db.Query("select id from user where username=? and password=?",username,password)

	if err != nil {
		fmt.Println("db.query is error login",err)
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println("failed to rows.Scan",err)
		}
		if id>0{
			message=`登录成功`
		}else{
			message=`用户名或密码不正确`
		}
	}
	return
}
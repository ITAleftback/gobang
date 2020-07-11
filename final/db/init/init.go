package init

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB

//连接数据库
func init()  {

	// 联机数据库 ==========================================================================
	Db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user?charset=utf8")

	err := Db.Ping()
	if err != nil {
		fmt.Println("fail to connect to db")
	}else {
		fmt.Println("连接数据库成功！")
	}
}

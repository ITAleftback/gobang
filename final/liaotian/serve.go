package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//加分项的玩家聊天
//====================================================
//先来建立一个客户端：
type Client struct {
	//用ID 来代替用户名

	//连接ws  方便用ws里的东西
	socket *websocket.Conn
	//想要发送的消息  存在切片中 并用通道发出去
	send chan []byte
}

//客户端管理
type ClientManager struct {
	//true 为在线  false为不在线 用map存储
	clients map[*Client]bool
	//用broadca来接受客户端发过来的消息
	broadcast chan []byte
	//登录上线
	register chan *Client
	//注销下线
	unregister chan *Client
}

//会把Message格式化成json
type Message struct {
	//消息struct
	Sender    string
	Recipient string
	Content   string
}

//创建客户端管理者  就是将前面的客户端管理结构体初始化 现在有了主人
var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:    make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

//然后写开始阶段 用方法 判断在线状态
// 这里应该与前端配合 只是拿来websocket测试的话用不上 但我还是根据网上的写下来 方便以后阅读
func (manager *ClientManager) start() {
	for {
		select {
		//如果有新的连接 通过channel传给conn
		case conn := <-manager.register:
			//设置为true 证明上线
			manager.clients[conn] = true
			//然后把连接成功的消息返回去
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected"})
			//调用客户端的send方法，发送消息给客户端
			manager.send(jsonMessage, conn)
		//如果连接断开
		case conn := <-manager.unregister:
			//判断bool  如果是true 就执行下面
			if _, ok := manager.clients[conn]; ok{
				//关闭send
				close(conn.send)
				//删除这个客户端
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		//广播
		case message:=<-manager.broadcast:
			//遍历已经连接的客户端，把消息发给他们
			for conn:=range manager.clients{
				conn.send<-message
			}

		}
	}
}

//写一个客户端的send方法
func (manager *ClientManager) send(message []byte, ignore *Client) {
	//遍历客户端 把消息发到每一个客户端上
	//这里要详细注释一下  ignore是传过来的新连接的客户端conn 然后遍历的时候  不能把消息发给自己  所以有！=语句
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
//来一个read的方法  这个读是客户来读 把放进web端的消息拿出来  并放入broadcast
func (c *Client)read(){
	//最后一定是要注销连接并关闭的
	defer func(){
		manager.unregister<-c
		c.socket.Close()
	}()
	for{
		//读取消息
		_,message,err:=c.socket.ReadMessage()
		//如果有错误信息就注销这个连接并关闭ws
		if err != nil {
			manager.unregister<-c
			_ = c.socket.Close()
			break
		}
		//没有的话就把信息放入broadcast
		jsonMessage,_:=json.Marshal(&Message{Content:string(message)})
		manager.broadcast<-jsonMessage
	}
}
//这个是从用户写在send的消息里读取并写入web端  上面的是从web端读
func (c *Client)write(){
	//defer func() {
	//	c.socket.Close()
	//}()
	//上面跟下面重复了，重复关闭了socket，经学长提醒已改正
	for{
		select {
		//从send里读取信息
		case message,_:=<-c.send:
			//如果没有消息
			//经学长提醒，将下列注释，如果进入了这个case说明一定是有消息的，不存在没有消息的情况
			//if !ok{
			//	_ = c.socket.WriteMessage(websocket.CloseMessage, []byte{})
			//	return
			//}
			//有消息就写入并发送到web端
			_ = c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}


//开一个路由
func main() {
	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":8080", nil)
}

//协议升级
func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	//这里有一个缺陷，我不知道该如何从登录状态拿到用户名
	//所以我随机生成了ID

	client := &Client{socket: conn, send: make(chan []byte)}
	manager.register <- client

	go client.read()
	go client.write()
}

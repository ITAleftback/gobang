
package main

import "log"


/*我看了了网上的斗地主的代码
但是还是无法独立写出五子棋的房间
在斗地主的基础上修改了一些
但不知道如何将进入房间 与路由之类联系起来
所以还是失败了
我先cope下来，日后若有能力也许会完善？
*/

// TPlayer 玩家类
type TPlayer struct {
	NickName string                  // 昵称
}

// TRoom 房间类
type TRoom struct {
}

// NewRoom 新建房间
func NewRoom() *TRoom {
	p := &TRoom{}

	return p
}

// QuickJoin 快速加入一张桌子
// pPlayer *TPlayer 玩家指针
func (self *TRoom) QuickJoin(pPlayer *TPlayer) bool {
	// 1 快速找到一个新桌子
	pTable := FindEmptyTable()
	if pTable == nil {
		log.Println("没有空桌子了. 需要新建一个空桌子")
		pTable = NewTable()
	}
	// 桌子里加入个新玩家
	pTable.playerJoin(pPlayer)
	// 如果桌子坐满了. 那么就开局
	if pTable.isFull() {
		pTable.playing()
	}

	return true
}

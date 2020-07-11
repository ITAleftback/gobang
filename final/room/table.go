
package main

import (
	"log"
)
//思想是想仿造qq游戏里面的桌子，但是能力有限没能实现
//此段代码是我从网上斗地主的桌子cope下来，也许日后能力提升会进行完善
var nTableAutoInc = 0    // 自增的桌子ID


// TTable 桌子类
type TTable struct {
	nIndex       int         // 桌子索引
	bFull        bool        // 是否满了
	pPlayerList  [2]*TPlayer // 里面的玩家
	nPlayerCount int

}

// NewTable 新建一个桌子
func NewTable() *TTable {
	nTableAutoInc++

	p := &TTable{}
	p.init()

	return p
}

// FindEmptyTable 找一个空的桌子
func FindEmptyTable() *TTable {
	var pTable *TTable
	return pTable
}

// 桌子开始玩耍
func (self *TTable) playing() {
	go func() {
		log.Println("游戏正式开始")

		//self.ddz = NewDDZ()
		log.Println("游戏正式结束")
	}()
}

// 假构造函数
func (self *TTable) init() {
	self.nPlayerCount = 0
}

// 桌子是否满了
func (self *TTable) isFull() bool {
	return self.nPlayerCount >= 3
}

// 桌子是否空的
func (self *TTable) isEmpty() bool {
	return self.nPlayerCount < 3
}

// 桌子上加入一个新玩家
func (self *TTable) playerJoin(pPlayer *TPlayer) {
	self.nPlayerCount++

	// 循环找空位插入
	for i := 0; i < 3; i++ {
		if self.pPlayerList[i] == nil {
			self.pPlayerList[i] = pPlayer
			return
		}
	}
}


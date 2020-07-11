package play

import (

	."final/game/chessboard"
	"final/game/judge"
	"fmt"
)
//这里就是下棋的过程了
func Play(i int32,j int32)(message string){
	if Board[i-1][j-1]!=0 {
		fmt.Println("这里已经有棋子了")
	}else {
		//玩家1执黑棋，第一个下
		if Flag%2 == 0 {
			Board[i-1][j-1] = 1
		} else {
			Board[i-1][j-1] = 2
		}

		for i := 0; i < 14; i++ {
			for j := 0; j < 14; j++ {
				fmt.Printf("%d\t", Board[i][j])
			}
			fmt.Println("\n")
		}
		//每次下完都要flag % 2 == 0进行判定是否输赢，若不分胜负则将次数++
		Flag++

		if judge.Judge(int(i-1), int(j-1)) {
			if Flag%2 == 0 {
				message=`玩家二胜利，游戏结束`
			} else {
				message=`玩家一胜利，游戏结束`
			}
		}
	}
	return
}
package chessboard

import "fmt"

/// 先来构造一个五子棋棋盘，因为无法做到可视化  所以用二维数组将就实现一些

var Flag =0 //表示下棋次数
var Board [15][15] int

func Chessboard() {
	//0代表无子，1代表黑子，2代表白子
	//i代表行   j代表列


	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			fmt.Printf("%d\t", Board[i][j])
		}
		fmt.Println("\n")
	}

}
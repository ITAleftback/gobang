package judge


import ."final/game/chessboard"
//判断五子棋输赢
func Judge(a int,b int)bool{
	var i,j int
	t := 2 - Flag% 2; //1 判断黑子是否赢	2 判断白子是否赢

	//纵向
	for i ,j =a - 4, b; i <= a; i++  {
		if i > 0 && i < 16 && t == Board[i][j] && t == Board[i + 1][j] && t == Board[i + 2][j] && t == Board[i + 3][j]  {
			return true;
		}
	}

	//横向
	for i , j =a, b - 4; j <= b; j++ {
		if (j > 0 && j < 16 && t == Board[i][j] && t == Board[i][j + 1] && t == Board[i][j + 2] && t == Board[i][j + 3] ) {
			return true;
		}
	}

	//右下
	for i, j = a-4,b - 4; i <= a&& j <= b; i,j=i+1,j+1 {
		if (i > 0 && i < 16 && j > 0 && j < 16 && t == Board[i][j] && t == Board[i + 1][j + 1] && t == Board[i + 2][j + 2] && t == Board[i + 3][j + 3] ) {
			return true;
		}
	}

	//左下
	for i , j = a - 4, b + 4; i <= a&& j >= b; i,j=i+1,j-1 {
		if (i > 0 && i < 16 && j > 0 && j < 16 && t == Board[i][j] && t == Board[i + 1][j - 1] && t == Board[i + 2][j - 2] && t == Board[i + 3][j - 3] ) {
			return true;
		}
	}

	return false
}
//投降用的，用来判断是哪个玩家选择投降根据Flag%2的值
func Sua()(message string){
	if Flag%2==0 {
		message=`游戏结束，玩家一认输`
	}else {
		message=`游戏结束，玩家二认输`
	}
	return
}
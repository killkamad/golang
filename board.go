package main
import (
	"fmt"
	"strings"
)

// Проверка есть ли число в массиве
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	var check []string
	var gameover bool
	var turn int = 1
	var playerItem string
	var x int
	var y int




	for gameover != true {
		//for i := range board {
		//	fmt.Println(board[i])
		//}
		if turn == 9 {
			gameover = true
			fmt.Println("Game over - ничья!")
		} else if turn % 2 == 0{
			fmt.Println("Ходит второй игрок - O")
			playerItem = "O"
		} else {
			fmt.Println("Ходит первый игрок - X")
			playerItem = "X"
		}

		fmt.Print("Введите x: ")
		fmt.Scan(&x)
		fmt.Print("Введите y: ")
		fmt.Scan(&y)
		value := fmt.Sprintf("%d%d", x,y)

		//proverka := contains(check, value)
		//fmt.Println("Занято место или нет = ",proverka)

		if (x <= 2 && x >= 0) && (y <= 2 && y >= 0){
			if contains(check, value) != true{
				board[x][y] = playerItem
				check = append(check, value)
				turn++
				//fmt.Printf("Циферки %s\n", check)
			} else {
				fmt.Println("*****МЕСТО ЗАНЯТО!*****")
			}

		} else {
			fmt.Println("ОШИБКА,\nИСПОЛЬЗУЙТЕ ЦИФРЫ от 0 - 2")
		}


		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	}
}

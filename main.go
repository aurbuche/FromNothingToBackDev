package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ( boardGame = [9]string{ // création d'un damier sans aucune cases remplies
"1", "2", "3",
"4", "5", "6",
"7", "8", "9"}
	endGame = false)

func PrintBoard() {
	for i := 0; i < 9; i++ {
		fmt.Print(" ", boardGame[i], " ")
		if (i+1) % 3 == 0 {
			fmt.Println()
		}
	}
}

func IsCaseAvailable(NoCase int, err error) bool {
	for i := 0; i < 9; i++ {
		if boardGame[NoCase] == "X" || boardGame[NoCase] == "0" {
			return false
		}
		i++
	}
	if NoCase < 1 || 9 < NoCase || err != nil{
		return false
	}
	return true
}

func AskPlayer(NoPlayer int) {
	var (
        NoCase  = 0
        err         error
        scanner     = bufio.NewScanner(os.Stdin)
    )
	fmt.Println("Enter a number between 1 and 9")
	scanner.Scan()
	NoCase, err = strconv.Atoi(scanner.Text())
	if IsCaseAvailable(NoCase - 1, err) {
		if NoPlayer == 1 {
			boardGame[NoCase - 1] = "X"
		} else {
			boardGame[NoCase - 1] = "O"
		}
	}
}

func main() {
	fmt.Println("Welcome to the tictactoe game:")
	for endGame == false {
		PrintBoard()
		AskPlayer(1);
		PrintBoard()
		AskPlayer(2);
	}
}

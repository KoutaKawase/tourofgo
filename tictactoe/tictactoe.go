package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	board := [][]string{
		[]string{"Y1", " _", " _", " _"},
		[]string{"Y2", " _", " _", " _"},
		[]string{"Y3", " _", " _", " _"},
		[]string{" ", "X1", "X2", "X3"},
	}
	inGame := true
	players := map[string]string{
		"circle": " 0",
		"cross":  " X",
	}
	current := players["circle"]
	reader := bufio.NewReader(os.Stdin)
	winner := ""

	clearConsole()

	for inGame {
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}

		fmt.Println("input your choice. Inputs are in X and Y order.")
		fmt.Println("Example) 2 3")
		fmt.Print(">> ")

		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		choiceBuffer := strings.Split(input, " ")
		choice := []int{}

		//入力された座標を配列のインデックス指定に使うために数値化
		for _, value := range choiceBuffer {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			choice = append(choice, num)
		}
		//仮に2が入力されたとして配列も2であって欲しいのでこのままでいい
		x := choice[0]
		//仮に2が入力されたらindexは1でなければならないので-1している
		y := choice[1] - 1

		if current == players["circle"] {
			board[y][x] = players["circle"]
			current = players["cross"]
		} else {
			board[y][x] = players["cross"]
			current = players["circle"]
		}

		winner = players["cross"]

		if winner != "" {
			inGame = false
		}
	}

	fmt.Printf("WINNER: %s", winner)
}

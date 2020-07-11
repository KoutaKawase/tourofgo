package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//GetWinner 現在のボードと現在のプレイヤーから勝者を出す
func GetWinner(board [][]string, current string) string {
	//-1してるのはX1..の部分を弾くため
	//横一列が揃ったパターンを走査
	for y := 0; y < len(board)-1; y++ {
		if board[y][1] == current && board[y][2] == current && board[y][3] == current {
			return current
		}
	}

	//縦のパターン走査
	//len(board[0])なのは横一列の長さを知りたかったので
	for x := 1; x < len(board[0]); x++ {
		if board[0][x] == current && board[1][x] == current && board[2][x] == current {
			return current
		}
	}

	//斜めのパターン走査
	//左上から右下への斜め走査
	if board[0][1] == current && board[1][2] == current && board[2][3] == current || board[0][3] == current && board[1][2] == current && board[2][1] == current {
		return current
	}

	return ""
}

//ReadConsole [1-3] [1-3] 以外の入力を受け付けないコンソール入力関数
func ReadConsole(reader *bufio.Reader) string {
	re := regexp.MustCompile("^[1-3] {1}[1-3]$")
	var input string
	isValidInput := false

	for !isValidInput {
		fmt.Print(">> ")
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if !re.MatchString(input) {
			fmt.Println("invalid input")
			continue
		}

		if existsAlready(input) {
			fmt.Println("The point is already placed.")
			continue
		}

		placedAlready = append(placedAlready, input)
		isValidInput = true
	}
	return input
}

func existsAlready(input string) bool {
	for _, value := range placedAlready {
		if value == input {
			return true
		}
	}
	return false
}

//Players プレイヤーの表記達
var Players map[string]string = map[string]string{
	"circle": " 0",
	"cross":  " X",
}

var placedAlready []string

func main() {
	board := [][]string{
		[]string{"Y1", " _", " _", " _"},
		[]string{"Y2", " _", " _", " _"},
		[]string{"Y3", " _", " _", " _"},
		[]string{" ", " X1", "X2", "X3"},
	}
	inGame := true
	current := Players["circle"]
	reader := bufio.NewReader(os.Stdin)
	winner := ""

	clearConsole()

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	for inGame {
		fmt.Println("input your choice. Inputs are in X and Y order.")
		fmt.Println("Example) 2 3")

		// input, _ := reader.ReadString('\n')
		input := ReadConsole(reader)

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

		if current == Players["circle"] {
			board[y][x] = Players["circle"]
			winner = GetWinner(board, current)
			current = Players["cross"]
		} else {
			board[y][x] = Players["cross"]
			winner = GetWinner(board, current)
			current = Players["circle"]
		}

		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}

		if winner != "" {
			inGame = false
		}
	}

	fmt.Printf("WINNER: %s", winner)
}

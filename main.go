package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	isWhite := false
	var board Board
	board.Init()
	board.Print()
	fmt.Println("Select your move like that: B 4")

	for {
		if isWhite {
			fmt.Print("White Turn:")
		} else {
			fmt.Print("Black Turn:")
		}
		input, _ := reader.ReadString('\n')
		b, x, y := CheckInput(input)
		if b == false || !board.Move(x, y, isWhite) {
			continue
		}
		if board.IsGameOver() {
			if isWhite {
				fmt.Println("White Won")
			} else {
				fmt.Println("Black Black")
			}
			return
		}
		isWhite = !isWhite
	}
}

func CheckInput(input string) (bool, int, int) {
	input = strings.ToUpper(input)
	l := strings.Split(input, " ")

	if len(l) != 2 {
		return false, 0, 0
	}
	x := int([]byte(l[0])[0]) - 65

	if x < 0 || x > 18 {
		return false, 0, 0
	}

	y, err := strconv.Atoi(strings.Split(l[1], "\n")[0])

	if err != nil || y < 0 || y > 19 {
		return false, 0, 0
	}
	return true, x, y - 1
}

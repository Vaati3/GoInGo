package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Empty rune = '🟨'
const Black rune = '⚫'
const White rune = '⚪'

type Board struct {
	height, width int
	tiles         [19][]rune
}

func (board *Board) Init() {
	board.height = 19
	board.width = 19
	for i := range board.height {
		var line strings.Builder

		for range board.width {
			line.WriteRune(Empty)
		}
		board.tiles[i] = []rune(line.String())
	}
}

func (board *Board) Print() {
	var letters strings.Builder
	letters.WriteString("   ")
	for j := range board.width {
		letters.WriteString(string(j+'A') + "  ")
	}
	fmt.Println(letters.String())
	for i := range board.height {
		var line strings.Builder
		if i < 10 {
			line.WriteString(" ")
		}
		line.WriteString(strconv.Itoa(i))
		line.WriteString(string(board.tiles[i]))
		fmt.Println(line.String())
	}
}

func (board *Board) Move(x, y int, isWhite bool) bool {
	if board.tiles[y][x] != Empty {
		return false
	}

	if isWhite {
		board.tiles[y][x] = White
	} else {
		board.tiles[y][x] = Black
	}
	board.Print()
	return true
}

func (board *Board) IsGameOver() bool {
	//TO DO

	return false
}

/*func (board *Board) GetNeighbors(x, y int) []rune {
	neighbors := []rune{}

}*/
